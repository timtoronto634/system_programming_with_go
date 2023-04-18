package chapter5httpserver

import (
	"bufio"
	"net"
	"net/http"
	"net/http/httputil"
)

// dial localhost:8080
func Dial() {
	conn, err := net.Dial("tcp", "localhost:8887")
	if err != nil {
		panic(err)
	}
	request, _ := http.NewRequest(
		"GET",
		"http://localhost:8887",
		nil,
	)
	err = request.Write(conn)
	if err != nil {
		panic(err)
	}
	response, _ := http.ReadResponse(bufio.NewReader(conn), request)
	dump, _ := httputil.DumpResponse(response, true)
	println(string(dump))
	conn.Close()
}
