package main

import (
	"log"
	"net"

	"golang.org/x/net/ipv4"
)

func main() {

	// ifconfigで確認可能
	en0, err := net.InterfaceByName("eth0")
	chkSe(err)
	// en1, err := net.InterfaceByIndex(65536)

	group := net.IPv4(224, 0, 0, 250)

	c, err := net.ListenPacket("udp4", "0.0.0.0:1024")
	chkSe(err)

	defer c.Close()

	p := ipv4.NewPacketConn(c)
	err = p.JoinGroup(en0, &net.UDPAddr{IP: group})
	chkSe(err)

	// err = p.JoinGroup(en1, &net.UDPAddr{IP: group})
	// chkSe(err)

	// message を有効にする
	err = p.SetControlMessage(ipv4.FlagDst, true)
	chkSe(err)

	b := make([]byte, 1500)

	log.Println("a")
	for {
		log.Println("b")
		n, cm, src, err := p.ReadFrom(b) // ここで待機状態になる
		log.Println("c")
		chkSe(err)

		if cm.Dst.IsMulticast() {
			go func() {
				log.Printf("Reciving data: %s from %s\n", string(b[:n]), src.String())
			}()

		} else {
			log.Println("is not multicast")
		}
	}

}

func chkSe(err error) {

	if err == nil {
		return
	}
	panic(err)
}
