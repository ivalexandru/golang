package main

import (
	"fmt"
)

//RETURNEAZA SI VALOAREA PT INT SI true/false
func half(n int) (int, bool) {
	return n / 2, n%2 == 0
}

func main() {
	h, even := half(5) //h e int, even este bool
	fmt.Println(h, even)
}
