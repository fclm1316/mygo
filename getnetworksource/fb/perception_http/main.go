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

var (
	handle *pcap.Handle
	err    error
	Client http.Client
)
var DataChan = make(chan []byte, 1000)

const version = "0.0.5"

func usage() {
	fmt.Printf("Usage of %s -i=\"eth0\" -f=\"tcp and port 8080\" -u=\"http://127.0.0.1/api\"\n", os.Args[0])
	fmt.Println("Options :")
	flag.PrintDefaults()
	fmt.Println("Version : ", version)
}

var url = flag.String("u", "http://127.0.0.1:12306/api", "send data to server api")
var iface = flag.String("i", "eth0", "Interface to get packets from")
var filter = flag.String("f", "tcp and port 8080", "BPF filter for pcap,usage: tcp and port 8080 or tcp ")
var snapshot_len = flag.Int64("s", 1024, "SnapLen for pcap packet capture,max 65535")
var debug = flag.Bool("debug", false, "debug flag, http://ip:2345/debug/pprof/")

func main() {
	flag.Usage = usage
	flag.Parse()
	log.Println("开始运行")

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
	go OpenDevice()
	go sendRequest()

	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM,
		syscall.SIGQUIT)
	func() {
		for s := range c {
			switch s {
			case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
				log.Println("close handle")
				defer handle.Close()
				defer close(DataChan)
				log.Println("exiting.....")
				os.Exit(0)
			default:
			}
		}
	}()
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
		go analyzePacketInfo(packet)
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

func analyzePacketInfo(packet gopacket.Packet) {
	var ipv4 *layers.IPv4
	var tcp *layers.TCP
	var eth *layers.Ethernet
	var payload []byte
	packettime := packet.Metadata().Timestamp
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
	applicationLayer := packet.ApplicationLayer()
	// 应用层
	if applicationLayer != nil {
		payload = applicationLayer.Payload()

		ethernetLayer := packet.Layer(layers.LayerTypeEthernet)
		eth, _ = ethernetLayer.(*layers.Ethernet)

		tcpLayer := packet.Layer(layers.LayerTypeTCP)
		tcp, _ = tcpLayer.(*layers.TCP)

		ipLayer := packet.Layer(layers.LayerTypeIPv4)
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
				request, err := Client.Post(*url, "application/json", bytes.NewBuffer(data))
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
