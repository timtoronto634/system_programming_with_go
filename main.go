package main

import (
	"os"

	ch8 "github.com/timtoronto634/system_programming_with_go/chapter8_unix_domain_socket"
)

func main() {
	args := os.Args[1:]
	if args[0] == "server" {
		ch8.Serve()
	} else {
		ch8.ClientDo()
	}
}
