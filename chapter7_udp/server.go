package chapter7udp

import (
	"fmt"
	"net"
)

func Serve() error {
	port := ":8887"
	// open socket
	conn, err := net.ListenPacket("udp", "localhost:8887")
	if err != nil {
		panic(err)
	}
	fmt.Printf("server is running at port%s\n", port)
	defer conn.Close()

	buffer := make([]byte, 1500)
	for {
		length, remoteAddress, err := conn.ReadFrom(buffer)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Received from %s: %v\n", remoteAddress, string(buffer[:length]))
		_, err = conn.WriteTo([]byte("Hello from Server"), remoteAddress)
		if err != nil {
			panic(err)
		}
	}
}
