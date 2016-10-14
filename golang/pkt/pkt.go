package main

import (
	"encoding/hex"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"log"
	"net"
)

func main() {

	buf := gopacket.NewSerializeBuffer()
	opts := gopacket.SerializeOptions{ComputeChecksums: true, FixLengths: true}

	eth := layers.Ethernet{
		DstMAC:       net.HardwareAddr{0x00, 0x0a, 0xf8, 0xAA, 0xBB, 0x44},
		SrcMAC:       net.HardwareAddr{0x00, 0x0a, 0xf8, 0xAA, 0xBB, 0x11},
		EthernetType: layers.EthernetTypeIPv4,
	}
	ip := layers.IPv4{
		Version:  4,
		IHL:      5,
		Id:       0x2222,
		SrcIP:    net.IP{0x011, 0x12, 0x13, 0x14},
		DstIP:    net.IP{0x21, 0x22, 0x23, 0x25},
		Length:   20 + 8 + 28,
		TTL:      2,
		Protocol: layers.IPProtocolUDP,
	}

	udp := layers.UDP{
		SrcPort:  0x11,
		DstPort:  0x53,
		Checksum: 0xffff,
		Length:   8 + 28,
	}

	ud := gopacket.Payload([]byte{
		0xDA, 0xDA, 0XDA, 0xDA,
		0xDA, 0xDA, 0xC1, 0x5C,
                0X00, 0xC1, 0x5C, 0X00,
                0xC1, 0x5C, 0x00, 0XC1,
                0x5C, 0x00, 0XEE, 0xEE,
                0xF0, 0x0D, 0XF0, 0x0D,
                0xF0, 0x0D, 0XF0, 0x0D,
	})

	udp.SetNetworkLayerForChecksum(&ip)

	if err := gopacket.SerializeLayers(buf, opts,
		&eth, &ip, &udp, ud); err != nil {
		log.Printf("failed to create pkt %s", err)
		return
	}

	packetData := buf.Bytes()

	log.Printf("generated pkt len: %d \n%s ", len(packetData), hex.Dump(packetData))

	packet := gopacket.NewPacket(packetData, layers.LayerTypeEthernet, gopacket.Default)
	// Get the UDP layer from this packet
	if udpLayer := packet.Layer(layers.LayerTypeUDP); udpLayer != nil {
		log.Printf("This is a UDP packet \n")
		// Get actual UDP data from this layer
		udp, _ := udpLayer.(*layers.UDP)
		log.Printf("From src port %d to dst port %d\n", udp.SrcPort, udp.DstPort)
	}
	// Iterate over all layers, printing out each layer type
	for _, layer := range packet.Layers() {
		log.Println("PACKET LAYER:", layer.LayerType())
	}
}
