package main

import "net"

func main() {

	en0, err := net.InterfaceByName("eth0")
	chkSe(err)

	group := net.IPv4(244, 0, 0, 255)

	c, err := net.ListenPackage("udp4")
}

func chkSe(err error) {

	if err == nil {
		return
	}

	panic(err)
}
