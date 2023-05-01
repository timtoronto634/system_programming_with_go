package chapter8domainsocket

import (
	"net"
)

func ClientDo() {
	conn, err := net.Dial("unix", "socketfile")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

}
