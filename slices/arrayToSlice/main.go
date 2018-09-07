package main

import (
	"crypto/rand"
	"fmt"
	"io"
)

func main() {
	var nonce [24]byte
	fmt.Println(nonce)
	//[0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
	// sunt 24 de zero
	io.ReadFull(rand.Reader, nonce[:]) //a transformat arrayul in slice
	//ReadFull e din pachetul io
	fmt.Println(nonce)
}
