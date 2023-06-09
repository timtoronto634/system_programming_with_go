package ch3

import (
	"crypto/rand"
	"io"
	"os"
)

const curDir = "chapter3_reader_interface/"

func Q3_1() {
	oldFile := curDir + "old.txt"
	newFile := curDir + "new.txt"
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

func Q3_2() {
	randReader := rand.Reader
	rf, err := os.Create(curDir + "rand.bin")
	if err != nil {
		panic(err)
	}
	defer rf.Close()
	io.CopyN(rf, randReader, 1024)
}
