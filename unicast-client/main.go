package main

import (
	"log"
	"net"
	"os"
)

func main() {

	conn, err := net.Dial("udp", "127.0.0.1:8080")
	chkSe(err)

	defer conn.Close()

	buf := []byte("Ping")
	n, err := conn.Write(buf)
	chkSe(err)

	if len(buf) != n {
		log.Printf("data size is %d, but sent data size is %d", len(buf), n)
	}

	recvBuf := make([]byte, 1024)
	n, err = conn.Read(recvBuf)
	chkSe(err)
	log.Printf("Received data: %s", string(recvBuf[:n]))
}

// chkSe .
func chkSe(err error) {
	if err == nil {
		return
	}
	log.Fatalln(err)
	os.Exit(1)
}
