package ch3

import (
	"io"
	"os"
)

func Q3_1() {
	oldFile := "chapter3_reader_interface/old.txt"
	newFile := "chapter3_reader_interface/new.txt"
	of, err := os.Open(oldFile)
	if err != nil {
		panic(err)
	}
	defer of.Close()
	nf, err := os.Create(newFile)
	if err != nil {
		panic(err)
	}
	io.Copy(nf, of)
	nf.Close()
}
