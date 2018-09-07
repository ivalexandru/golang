package main

import "fmt"

func main() {
	fmt.Println([]byte("Hello")) //[72 101 108 108 111] *1
}

// *1  converts this string of characters into a SLICE of bytes
// 72, de exemplu este H in tabelul UNICODE - ascii.
