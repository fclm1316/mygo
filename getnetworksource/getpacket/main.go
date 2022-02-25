package main

import (
	// "bytes"
	// "encoding/json"
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
	Th_dport uint16 `json:"th_dport"`
	Th_sport uint16 `json:"th_sport"`

	Timestamp int64  `json:"timestamp"`
	Ip_dst    string `json:"ip_dst"`
	Ip_src    string `json:"ip_src"`

	Body string `json:"body"`
}

var (
	handle *pcap.Handle
	err    error
)

func usage() {
	fmt.Printf("Usage of %s -i=\"eth0\" -f=\"tcp and port 8080\" -u=\"http://127.0.0.1/api\"\n", os.Args[0])
	fmt.Println("Options :")
	flag.PrintDefaults()
}

var url = flag.String("u", "http://127.0.0.1:12306/api", "send data to server api")
var iface = flag.String("i", "eth0", "Interface to get packets from")
var filter = flag.String("f", "tcp and port 8080", "BPF filter for pcap,usage: tcp and port 8080 or tcp ")
var snapshot_len = flag.Int64("s", 1024, "SnapLen for pcap packet capture,max 65535")

func main() {
	flag.Usage = usage
	flag.Parse()
	log.Println("开始运行")

	go OpenDevice()

	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM,
		syscall.SIGQUIT)
	func() {
		for s := range c {
			switch s {
			case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
				log.Println("close handle")
				defer handle.Close()
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

		tcpLayer := packet.Layer(layers.LayerTypeTCP)
		tcp, _ = tcpLayer.(*layers.TCP)

		ipLayer := packet.Layer(layers.LayerTypeIPv4)
		ipv4, _ = ipLayer.(*layers.IPv4)

		ReqData := NewRequestData(uint16(tcp.DstPort), uint16(tcp.SrcPort), packettime.UnixNano(), ipv4.DstIP.String(), ipv4.SrcIP.String(), string(payload))

		log.Println("==================")
		fmt.Printf("%v \n", ReqData.Timestamp)
		fmt.Println("srcip: %s,srcport: %v\n", ReqData.Ip_src, ReqData.Th_sport)
		fmt.Println("dstip: %s,desport: %v\n", ReqData.Ip_dst, ReqData.Th_dport)
		fmt.Println("body：%v", ReqData.Body)
		log.Println("------------------")
	}
}

func NewRequestData(th_dport uint16, th_sport uint16, timestamp int64, ip_dst string, ip_src string, body string) *RequestData {
	return &RequestData{

		Th_dport: th_dport,
		Th_sport: th_sport,

		Timestamp: timestamp,
		Ip_dst:    ip_dst,
		Ip_src:    ip_src,

		Body: body,
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
