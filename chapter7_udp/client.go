package chapter7udp

import (
	"fmt"
	"net"
)

func ClientDo() {
	conn, err := net.Dial("udp", "localhost:8887")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	fmt.Println("Sending to server")
	_, err = conn.Write([]byte("Hello from Client"))
	if err != nil {
		panic(err)
	}

	fmt.Println("receiving from server")
	buffer := make([]byte, 1500)
	length, err := conn.Read(buffer)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Received: %s\n", string(buffer[:length]))
}
