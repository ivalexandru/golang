package main

import "fmt"

func main() {

	myGreeting := map[int]string{
		0: "zero",
		1: "bnjr",
		2: "buenos dias",
		3: "buongiorno",
	}

	for cheie, valoare := range myGreeting {
		fmt.Println(cheie, " - ", valoare)
	}
}
