package main

import (
	"os"

	ch6 "github.com/timtoronto634/system_programming_with_go/chapter6_http_server"
)

func main() {
	args := os.Args[1:]
	if args[0] == "server" {
		ch6.Serve()
	} else {
		ch6.ClientDo()
	}
}
