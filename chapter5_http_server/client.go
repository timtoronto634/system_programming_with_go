package chapter5httpserver

import "net"

// dial localhost:8080
func Dial() {
	conn, err := net.Dial("tcp", "localhost:8887")
	if err != nil {
		panic(err)
	}
	conn.Close()
}
