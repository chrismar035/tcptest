package main

import (
	"net"
	"os"
)

func main() {
	ip, port := "192.168.1.69", "333"

	tvAddr, err := net.ResolveTCPAddr("tcp", ip+":"+port)
	if err != nil {
		println("Could not resolve", ip, "on", port)
		os.Exit(1)
	}

	conn, err := net.DialTCP("tcp", nil, tvAddr)
	if err != nil {
		println("Could not connect to", ip, "on", port)
		os.Exit(1)
	}

	var recv []byte
	conn.Read(recv)
	conn.Write([]byte("led-on\r\n"))
}
