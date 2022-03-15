package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding/simplifiedchinese"
)

type RequestData struct {
	Ack       bool   `json:"ack"`
	Fin       bool   `json:"fin"`
	Syn       bool   `json:"syn"`
	Th_dport  uint16 `json:"th_dport"`
	Th_sport  uint16 `json:"th_sport"`
	Lens      uint16 `json:"lens"`
	Seq       uint32 `json:"seq"`
	Ack_seq   uint32 `json:"ack_seq"`
	Timestamp int64  `json:"timestamp"`
	Ip_dst    string `json:"ip_dst"`
	Ip_src    string `json:"ip_src"`
	Mac_dst   string `json:"mac_dst"`
	Mac_src   string `json:"mac_src"`
	Body      string `json:"body"`
}

var heartbeat struct {
	period time.Duration
}

var (
	handle *pcap.Handle
	err    error
	Client http.Client
)
var DataChan = make(chan []byte, 30)
var PacaetChan = make(chan interface{}, 30)

const (
	version         = "0.0.6"
	reciveTcpPacket = "api/reciveTcpPacket"
	register        = "api/register"
)

func usage() {
	fmt.Printf("Usage of %s -i=\"eth0\" -f=\"tcp and port 8080\" -s=\"127.0.0.1:8080\"\n", os.Args[0])
	fmt.Println("Options :")
	flag.PrintDefaults()
	fmt.Println("Version : ", version)
}

var serverip = flag.String("s", "127.0.0.1:12306", "send data to server api")
var iface = flag.String("i", "eth0", "Interface to get packets from")
var filter = flag.String("f", "tcp and port 8080", "BPF filter for pcap,usage: tcp and port 8080 or tcp ")
var snapshot_len = flag.Int64("l", 1024, "SnapLen for pcap packet capture,max 65535")
var debug = flag.Bool("debug", false, "debug flag, http://ip:2345/debug/pprof/")
var name = flag.String("n", "default", "application name")
var env = flag.String("e", "FAT", "env: DEV,FAT,UAT,LPT,PRO")

func main() {
	flag.Usage = usage
	flag.Parse()
	log.Println("开始运行")
	var urlRT = fmt.Sprintf("http://%s/%s", *serverip, reciveTcpPacket)
	var urlRG = fmt.Sprintf("http://%s/%s", *serverip, register)
	log.Println("debug : ", *debug)
	t := http.DefaultTransport.(*http.Transport).Clone()
	t.MaxIdleConns = 30
	t.MaxConnsPerHost = 30
	t.MaxIdleConnsPerHost = 30
	Client = http.Client{
		Timeout:   10 * time.Second,
		Transport: t,
	}
	if *debug {
		go func() {
			http.ListenAndServe("0.0.0.0:2345", nil)
		}()
	}
	defer handle.Close()
	defer close(DataChan)
	defer close(PacaetChan)

	go sendRequest()
	go analyzePacketInfo()
	go heartbeatNew(time.Minute).Start(urlRG)
	OpenDevice()
	// c := make(chan os.Signal)
	// signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM,
	// 	syscall.SIGQUIT)
	// func() {
	// 	for s := range c {
	// 		switch s {
	// 		case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
	// 			log.Println("close handle")
	// 			defer handle.Close()
	// 			defer close(DataChan)
	// 			log.Println("exiting.....")
	// 			os.Exit(0)
	// 		default:
	// 		}
	// 	}
	// }()
}

func heartbeatNew(t time.Duration) *heartbeat {
	return &heartbeat{
		period: t,
	}
}

// 增加心跳
func (beat *heartbeat) Start(url string) *heartbeat {
	ticker := time.Ticker(beat.period)
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	q := req.URL.Query()
	q.Add("appName", *name)
	q.Add("env", *env)
	req.URL.RawQuery = q.Encode()
	for range ticker.C {
		res, err := Client.Do(req)
		if err != nil {
			log.Println(err)
			continue
		}
		res.Body.Close()
	}
	return nil
}
func OpenDevice() {
	if *snapshot_len > 65535 || *snapshot_len < 0 {
		log.Println("Check snapshot_len")
		log.Fatal(*snapshot_len)
	}
	handle, err = pcap.OpenLive(*iface, int32(*snapshot_len), true, pcap.BlockForever)
	if err != nil {
		log.Fatal(err)
	}
	if err := handle.SetBPFFilter(*filter); err != nil {
		log.Fatal(err)
	}
	log.Printf("打开设备: %v", *iface)
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	packetSource.NoCopy = true
	packets := packetSource.Packets()

	for {
		packet := <-packets
		if packet == nil {
			return
		}
		if packet.NetworkLayer() == nil || packet.TransportLayer() == nil || packet.TransportLayer().LayerType() != layers.LayerTypeTCP {
			log.Println("Unusable packet")
			continue
		}
		// go analyzePacketInfo(packet)
		PacaetChan <- packet
	}
}
func stopSelf(handle *pcap.Handle) {
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM,
		syscall.SIGQUIT)
	func() {
		for s := range c {
			switch s {
			case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
				log.Println("close handle")
				handle.Close()
				log.Println("exiting.....")
				os.Exit(0)
			default:
			}
		}
	}()
}

// 编码转换
func covCharset(b []byte) []beat {
	_, name, _ := charset.DetermineEncoding(b, "text/html")
	if name != "utf-8" {
		decodeByte, _ := simplifiedchinese.GB18030.NewDecoder().Bytes(b)
		return decodeByte
	}
	return b
}

func analyzePacketInfo(packet gopacket.Packet) {
	for {
		select {
		case packet := PacaetChan:
			go func() {
				var ipv4 *layers.IPv4
				var tcp *layers.TCP
				var eth *layers.Ethernet
				var payload []byte
				// 空接口转普通接口 type interface{} is interface with no methods
				packettime := packet.(gopacket.Packet).Metadata().Timestamp
				//ethernetLayer := packet.Layer(layers.LayerTypeEthernet)
				//if ethernetLayer != nil {
				//    eth, _ = ethernetLayer.(*layers.Ethernet)
				//}
				//tcpLayer := packet.Layer(layers.LayerTypeTCP)
				//// tcp
				//if tcpLayer != nil {
				//    tcp, _ = tcpLayer.(*layers.TCP)
				//}
				//ipLayer := packet.Layer(layers.LayerTypeIPv4)
				//// IPv4
				//if ipLayer != nil {
				//    ipv4, _ = ipLayer.(*layers.IPv4)
				//}
				applicationLayer := packet.(gopacket.Packet).ApplicationLayer()
				// 应用层
				if applicationLayer != nil {
					payload = covCharset(applicationLayer.Payload())

					ethernetLayer := packet.(gopacket.Packet).Layer(layers.LayerTypeEthernet)
					eth, _ = ethernetLayer.(*layers.Ethernet)

					tcpLayer := packet.(gopacket.Packet).Layer(layers.LayerTypeTCP)
					tcp, _ = tcpLayer.(*layers.TCP)

					ipLayer := packet.(gopacket.Packet).Layer(layers.LayerTypeIPv4)
					ipv4, _ = ipLayer.(*layers.IPv4)

					ReqData := NewRequestData(tcp.ACK, tcp.FIN, tcp.SYN, uint16(tcp.DstPort), uint16(tcp.SrcPort), ipv4.Length, tcp.Seq,
						tcp.Ack, packettime.UnixNano(), ipv4.DstIP.String(), ipv4.SrcIP.String(), eth.DstMAC.String(), eth.SrcMAC.String(), string(payload))

					data, err := json.Marshal(ReqData)
					if err != nil {
						log.Println(err)
					}
					// go sendRequest(data)
					DataChan <- data
				}
			}()
		default:
			time.Sleep(1e9)
		}
	}
}

func NewRequestData(ack bool, fin bool, syn bool, th_dport uint16, th_sport uint16, lens uint16, seq uint32,
	ack_seq uint32, timestamp int64, ip_dst string, ip_src string, mac_dst string, mac_src string, body string) *RequestData {
	return &RequestData{
		Ack:       ack,
		Fin:       fin,
		Syn:       syn,
		Th_dport:  th_dport,
		Th_sport:  th_sport,
		Lens:      lens,
		Seq:       seq,
		Ack_seq:   ack_seq,
		Timestamp: timestamp,
		Ip_dst:    ip_dst,
		Ip_src:    ip_src,
		Mac_dst:   mac_dst,
		Mac_src:   mac_src,
		Body:      body,
	}
}

func sendRequest() {
	// 发送数据
	for {
		select {
		case data := <-DataChan:
			go func(d []byte) {
				request, err := Client.Post(*url, "application/json,charset=UTF-8", bytes.NewBuffer(data))
				if err != nil {
					log.Println(err)
					break
				}
				if *debug {
					log.Printf(string(data))
					body, _ := ioutil.ReadAll(request.Body)
					log.Println(string(data))
				}
				request.Body.Close()
			}(data)

		default:
			time.Sleep(1e9)
		}
	}
}

// func sendRequest(data []byte) {
// 	//log.Println(string(data))

// 	request, err := Client.Post(*url, "application/json", bytes.NewReader(data))
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}
// 	defer request.Body.Close()
// 	//body,_ := ioutil.ReadAll(request.Body)
// 	//log.Println(string(body))
// }
