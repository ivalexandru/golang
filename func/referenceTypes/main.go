package main

import "fmt"

func main() {
	m := make([]string, 1, 25)
	fmt.Println(m) // slice gol  []
	changeMe(m)
	fmt.Println(m) //TODD si aici => PASS BY VALUE
}

func changeMe(z []string) {
	z[0] = "Todd"
	fmt.Println(z) //Todd
}
