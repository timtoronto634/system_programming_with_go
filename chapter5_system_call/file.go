package ch5

import (
	"os"
	"strings"
)

func CreateFile(filename string) {
	if !strings.HasSuffix(filename, ".txt") {
		filename = filename + ".txt"
	}
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	file.WriteString("Hello World")
}
