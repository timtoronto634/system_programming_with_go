package chapter5httpserver

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
	"time"
)

func writeToConn(sessionResponses chan chan *http.Response, conn net.Conn) {
	defer conn.Close()

	for sessionResponse := range sessionResponses {
		response := <-sessionResponse
		response.Write(conn)
		close(sessionResponse)
	}
}

func handleRequest(request *http.Request, resultReceiver chan *http.Response) {
	dump, err := httputil.DumpRequest(request, true)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(dump))
	content := "Hello World\n"

	response := &http.Response{
		StatusCode:    200,
		ProtoMajor:    1,
		ProtoMinor:    1,
		ContentLength: int64(len(content)),
		Body:          io.NopCloser(strings.NewReader(content)),
	}

	resultReceiver <- response
}

func processSession(conn net.Conn) {
	defer conn.Close()
	fmt.Printf("Accept %v\n", conn.RemoteAddr())
	// channel to handle request inside session
	sessionResponses := make(chan chan *http.Response, 50)
	defer close(sessionResponses)
	// goroutine for serialize response
	go writeToConn(sessionResponses, conn)
	reader := bufio.NewReader(conn)
	for {
		conn.SetReadDeadline(time.Now().Add(5 * time.Second))
		request, err := http.ReadRequest(reader)
		if err != nil {
			// if  timeout, close connection,
			neterr, ok := err.(net.Error)
			if ok && neterr.Timeout() {
				fmt.Println("Timeout")
				break
			} else if err == io.EOF {
				break
			}
			panic(err)
		}
		sessionResponse := make(chan *http.Response)
		// sessionResponse は同期的にsessionResponsesに入れる
		sessionResponses <- sessionResponse

		// ただ単にhandleRequest の結果をchannel にいれるだけだと、他のリクエストが先に処理を終えて、channelに入ってしまう
		go handleRequest(request, sessionResponse)
	}
}

func Serve() error {
	port := ":8887"
	// open socket
	listener, err := net.Listen("tcp", port)
	if err != nil {
		panic(err)
	}
	fmt.Printf("server is running at port%s\n", port)
	for {
		// tcp connection
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		go processSession(conn)
	}
}
