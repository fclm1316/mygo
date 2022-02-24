package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"log"
	"regexp"
	//"golang.org/x/crypto/openpgp/packet"
	"io/ioutil"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

var requestChan = make(chan requestData, 1000)

type requestData struct {
	TimeNow      string
	EthernetType string
	Protocol     string
	SrcIp        string
	SrcPort      string
	DstIp        string
	DstPort      string
	Url          string
}

var (
	promisc      bool = false
	debug        bool
	snapshot_len int32 = 1024
	//timeout      time.Duration = 30 * time.Second
	timeout time.Duration = -1 * time.Second
	err     error
	handle  *pcap.Handle
	url     string
)

func init() {
	flag.StringVar(&url, "url", "http://127.0.0.1:12306/api", "输入数据接收端的地址")
	flag.BoolVar(&debug, "debug", false, "debug开启标志,访问 http://ip:2345/debug/pprof/")

}
func main() {
	flag.Parse()
	if debug {
		go func() {
			http.ListenAndServe("0.0.0.0:2345", nil)
		}()
	}
	//defer profile.Start(profile.MemProfile).Stop()
	//defer profile.Start(profile.CPUProfile).Stop()
	devName := findDevName()
	log.Println(devName)
	go openDevice(devName)
	comsterRequestData()
}
func stopSelf() {
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM,
		syscall.SIGQUIT)
	func() {
		for s := range c {
			switch s {
			case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
				log.Println("close chan")
				close(requestChan)
				os.Exit(0)
			default:
			}
		}
	}()
}
func findDevName() (deviceName string) {
	//fmt.Println("find dev")
	//var deviceName string
	devices, _ := pcap.FindAllDevs()
	for _, device := range devices {
		//fmt.Println(device.Addresses)
		//fmt.Println(len(device.Addresses))
		if len(device.Addresses) != 0 {
			deviceIP := device.Addresses[0].IP.String()
			//fmt.Println(reflect.TypeOf(deviceIP))
			if strings.HasPrefix(deviceIP, "172") || strings.HasPrefix(deviceIP, "10") {
				//fmt.Println(deviceIP.String())
				deviceName = device.Name
				break
			} else {
				deviceName = "ens192"
			}

		}
	}
	return deviceName
}

func openDevice(deviceName string) {
	//fmt.Println("open dev")
	handle, err = pcap.OpenLive(deviceName, snapshot_len, promisc, timeout)
	if err != nil {
		log.Printf("pcap open live failed %v\n", err)
		os.Exit(1)
	}
	//var filter string = "tcp"
	//err := handle.SetBPFFilter(filter)
	//if err != nil {
	//    fmt.Println("filter err")
	//}
	defer handle.Close()
	go stopSelf()
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	packetSource.NoCopy = true
	for packet := range packetSource.Packets() {
		go analyzePacketInfo(packet)
	}
}

func listAllLayer(packet gopacket.Packet) {
	for _, layer := range packet.Layers() {
		log.Println("- ", layer.LayerType())
	}
}

func analyzePacketInfo(packet gopacket.Packet) {
	var (
		SrcIP        string
		SrcPort      string
		DstIP        string
		DstPort      string
		EthernetType string
		Protocol     string
		ReqUrl       string
	)
	packettime := packet.Metadata().Timestamp.Format("2006-01-02 03:04:05")

	ethernetLayer := packet.Layer(layers.LayerTypeEthernet)
	//以太层
	if ethernetLayer != nil {
		//fmt.Println("Ethernet layer detected.")
		ethernetPacket, _ := ethernetLayer.(*layers.Ethernet)
		//fmt.Println("Source MAC: ", ethernetPacket.SrcMAC)
		//fmt.Println("Destination MAC: ", ethernetPacket.DstMAC)
		//fmt.Println("Ethernet type: ", ethernetPacket.EthernetType)
		EthernetType = ethernetPacket.EthernetType.String()
		//fmt.Println()
	}

	ipLayerV6 := packet.Layer(layers.LayerTypeIPv6)
	if ipLayerV6 != nil {
		// IPv6
		ipv6, _ := ipLayerV6.(*layers.IPv6)
		Protocol = ipv6.NextHeader.String()
		SrcIP = ipv6.SrcIP.String()
		DstIP = ipv6.DstIP.String()
	} else {
		Protocol = ""
		SrcIP = ""
		DstIP = ""
	}

	ipLayer := packet.Layer(layers.LayerTypeIPv4)
	// IPv4
	if ipLayer != nil {
		ipv4, _ := ipLayer.(*layers.IPv4)
		Protocol = ipv4.Protocol.String()
		SrcIP = ipv4.SrcIP.String()
		DstIP = ipv4.DstIP.String()
	} else {
		Protocol = ""
		SrcIP = ""
		DstIP = ""
	}

	tcpLayer := packet.Layer(layers.LayerTypeTCP)
	// TCP层
	if tcpLayer != nil {
		//fmt.Println("TCP layer detected.")
		tcp, _ := tcpLayer.(*layers.TCP)
		// TCP layer variables:
		// SrcPort, DstPort, Seq, Ack, DataOffset, Window, Checksum, Urgent
		// Bool flags: FIN, SYN, RST, PSH, ACK, URG, ECE, CWR, NS
		SrcPort = tcp.SrcPort.String()
		DstPort = tcp.DstPort.String()
	} else {
		SrcPort = ""
		DstPort = ""
	}

	applicationLayer := packet.ApplicationLayer()
	// 应用层
	if applicationLayer != nil {
		payload := string(applicationLayer.Payload())
		if strings.HasPrefix(payload, "GET") || strings.HasPrefix(payload, "POST") {
			reg := regexp.MustCompile(`(?s)(GET|POST) (.*?) HTTP.*Host: (.*?)\n`)
			result := reg.FindStringSubmatch(payload)
			var build strings.Builder
			build.WriteString("http://")
			build.WriteString(strings.TrimSpace(result[3]))
			build.WriteString(strings.TrimSpace(result[2]))
			ReqUrl = build.String()
			//ReqUrl = "http://" + strings.TrimSpace(result[3]) + strings.TrimSpace(result[2])

		} else {
			ReqUrl = ""
		}
	}

	// Check for errors
	if err := packet.ErrorLayer(); err != nil {
		log.Println("Error decoding some part of the packet:", err)
	}

	//var ip *layers.IPv4
	//var tcp *layers.TCP
	//var ReqUrl string
	// packettime := *sourcePacket
	//packettime := packet.Metadata().Timestamp.Format("2006-01-02 03:04:05")

	//ethernetLayer := packet.Layer(layers.LayerTypeEthernet)
	//if ethernetLayer != nil {
	//	ethernetPacket, _ := ethernetLayer.(*layers.Ethernet)
	//	// 过滤 IPv6
	//	if ethernetPacket.EthernetType.String() == "IPv4" {
	//		ipLayer := packet.Layer(layers.LayerTypeIPv4)
	//		if ipLayer != nil {
	//			// 获得地址
	//			ip, _ = ipLayer.(*layers.IPv4)
	//			tcpLayer := packet.Layer(layers.LayerTypeTCP)
	//			if tcpLayer != nil {
	//				// 获得端口
	//				tcp, _ = tcpLayer.(*layers.TCP)
	//				applicationLayer := packet.ApplicationLayer()
	//				// 解码应用层
	//				if applicationLayer != nil {
	//					payload := string(applicationLayer.Payload())
	//					//fmt.Println(payload)
	//					if strings.HasPrefix(payload, "GET") || strings.HasPrefix(payload, "POST") {
	//						reg := regexp.MustCompile(`(?s)(GET|POST) (.*?) HTTP.*Host: (.*?)\n`)
	//						result := reg.FindStringSubmatch(payload)
	//						ReqUrl = "http://" + strings.TrimSpace(result[3]) + strings.TrimSpace(result[2])
	//					}
	//				}
	//			}
	//		}
	//	}
	//}
	factoryRequestData(&packettime, &EthernetType, &Protocol, &SrcIP, &SrcPort, &DstIP, &DstPort, &ReqUrl)

}
func factoryRequestData(packtime *string, ethernetType *string, protocol *string, srcip *string, srcport *string, dstip *string, dstport *string, url *string) {
	//fmt.Println("in chan")
	NewData := requestData{
		TimeNow:      *packtime,
		EthernetType: *ethernetType,
		Protocol:     *protocol,
		SrcIp:        *srcip,
		SrcPort:      *srcport,
		DstIp:        *dstip,
		DstPort:      *dstport,
		Url:          *url,
	}
	requestChan <- NewData
}

func comsterRequestData() {
	for {
		if len(requestChan) != 0 {
			//fmt.Println("out chan")
			cc := <-requestChan
			//fmt.Println(cc)
			data, err := json.Marshal(cc)
			if err != nil {
				log.Println("json load failed")
				continue
			}
			//fmt.Println(string(data))
			go sendRequest(data)

		} else {
			//fmt.Println("sleep 5")
			time.Sleep(5 * 1e9)
		}
	}
}

func sendRequest(data []byte) {
	//url := "http://127.0.0.1:12306/api"
	request, err := http.Post(url, "application/json", bytes.NewReader(data))
	if err != nil {
		log.Println(err)
	}
	defer request.Body.Close()
	body, _ := ioutil.ReadAll(request.Body)
	log.Println(string(body))
}
