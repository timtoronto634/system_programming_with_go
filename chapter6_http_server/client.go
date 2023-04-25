package chapter5httpserver

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
)

// dial localhost:8080
func ClientDo() {
	sendMessages := []string{
		"ASCII",
		"PROGRAMMING",
		"PLUS",
	}
	var conn net.Conn = nil
	var err error
	requests := make([]*http.Request, 0, len(sendMessages))

	conn, err = net.Dial("tcp", "localhost:8887")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Access\n")
	defer conn.Close()

	for i := 0; i < len(sendMessages); i++ {
		lastMessage := i == len(sendMessages)-1
		request, err := http.NewRequest(
			"GET",
			"http://localhost:8887?message="+sendMessages[i],
			nil,
		)
		if lastMessage {
			request.Header.Add("Connection", "close")
		} else {
			request.Header.Add("Connection", "keep-alive")
		}
		if err != nil {
			panic(err)
		}
		err = request.Write(conn)
		if err != nil {
			panic(err)
		}
		fmt.Println("send: ", sendMessages[i])
		requests = append(requests, request)
	}

	reader := bufio.NewReader(conn)
	for _, request := range requests {
		response, err := http.ReadResponse(reader, request)
		if err != nil {
			panic(err)
		}
		dump, err := httputil.DumpResponse(response, false)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(dump))
	}
}
