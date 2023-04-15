package ch5

import (
	"os"
	"strings"
)

func CreateFile(filename string) {
	if !strings.HasSuffix(filename, ".txt") {
		filename = filename + ".txt"
	}
	// Create call will reach system call defined in C at: /go/src/runtime/internal/syscall/asm_linux_amd64.s
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	file.WriteString("Hello World")
}
