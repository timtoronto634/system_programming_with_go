package chapter8domainsocket

import (
	"net"
	"os"
)

func Serve() error {
	os.Remove("socketfile")
	listener, err := net.Listen("unix", "socketfile")
	if err != nil {
		panic(err)
	}
	defer listener.Close()
	_, err = listener.Accept()
	if err != nil {
		panic(err)
	}
	return nil
}
