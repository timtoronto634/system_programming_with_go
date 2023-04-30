package main

import (
	"os"

	ch7 "github.com/timtoronto634/system_programming_with_go/chapter7_udp"
)

func main() {
	args := os.Args[1:]
	if args[0] == "server" {
		ch7.Serve()
	} else {
		ch7.ClientDo()
	}
}
