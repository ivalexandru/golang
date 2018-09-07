package main

import (
	"fmt"
)

func wrapper() func() int { //unde func() int este TYPE-ul

	var x int //echiv cu    x := 0

	return func() int {
		x++
		return x
	}

}

func main() {
	increment := wrapper()
	fmt.Println(increment()) // 1  //*1
	fmt.Println(increment()) // 2
}

//*1  increment are () pt ca wrapper este de TYPE func() int, nu pt ca increment ar fi functie
