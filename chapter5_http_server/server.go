package chapter5httpserver

import (
	"fmt"
	"net"
)

func Serve() error {
	ln, err := net.Listen("tcp", ":8887")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}
		go func() {
			var b []byte
			_, err = conn.Read(b)
			if err != nil {
				panic(err)
			}

			fmt.Println(b)
		}()
	}
}
