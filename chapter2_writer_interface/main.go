package main

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Encoding", "gzip")
	w.Header().Set("Content-Type", "application/json")
	source := map[string]string{
		"text": "hello world",
	}

	multiWriter := io.MultiWriter(os.Stdout, w)

	gwriter := gzip.NewWriter(multiWriter)
	encoder := json.NewEncoder(gwriter)
	encoder.Encode(source)
	gwriter.Flush()
	gwriter.Close()
}

func q2_3() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func q2_1() {
	printTest, err := os.Create("print_test.txt")
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(printTest, "digit: %d, float %f, string: %s\n", 1, 1., "1")
}

func main() {
	q2_1()
	q2_3()
}
