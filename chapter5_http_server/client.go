package chapter5httpserver

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
)

// dial localhost:8080
func ClientDo() {
	sendMessages := []string{
		"ASCII",
		"PROGRAMMING",
		"PLUS",
	}
	current := 0
	var conn net.Conn = nil

	for {
		var err error
		if conn == nil {
			conn, err = net.Dial("tcp", "localhost:8887")
			if err != nil {
				panic(err)
			}
			fmt.Printf("Access: %d\n", current)
		}
		request, err := http.NewRequest(
			"POST",
			"http://localhost:8887",
			strings.NewReader(sendMessages[current]),
		)
		if err != nil {
			panic(err)
		}
		err = request.Write(conn)
		if err != nil {
			panic(err)
		}
		response, err := http.ReadResponse(
			bufio.NewReader(conn), request)
		if err != nil {
			fmt.Println("Retry")
			conn = nil
			continue
		}
		dump, err := httputil.DumpResponse(response, true)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(dump))
		current++

		if current == len(sendMessages) {
			break
		}
	}
}
