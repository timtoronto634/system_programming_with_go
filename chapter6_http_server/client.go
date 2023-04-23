package chapter5httpserver

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httputil"
	"strconv"
)

func ClientDo() {
	conn, err := net.Dial("tcp", "localhost:8887")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	request, err := http.NewRequest(
		"GET",
		"http://localhost:8887",
		nil,
	)
	if err != nil {
		panic(err)
	}
	err = request.Write(conn)
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(conn)
	response, err := http.ReadResponse(
		reader, request)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	dump, err := httputil.DumpResponse(response, false)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(dump))
	if len(response.TransferEncoding) < 1 ||
		response.TransferEncoding[0] != "chunked" {
		panic("wrong transfer encoding")
	}
	for {
		sizeStr, err := reader.ReadBytes('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
		// parse size, expressed in hex, close if the size is zero
		size, err := strconv.ParseInt(string(sizeStr[:len(sizeStr)-2]), 16, 64)
		if err != nil {
			panic(err)
		}
		if size == 0 {
			break
		}
		if err != nil {
			panic(err)
		}
		line := make([]byte, int(size))
		io.ReadFull(reader, line)
		reader.Discard(2)
		fmt.Printf("  %d bytes: %s\n", size, string(line))
	}
}
