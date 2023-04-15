package main

import (
	"os"

	ch5 "github.com/timtoronto634/system_programming_with_go/chapter5_http_server"
)

func main() {
	args := os.Args[1:]
	if args[0] == "server" {
		ch5.Serve()
	} else {
		ch5.Dial()
	}
}
