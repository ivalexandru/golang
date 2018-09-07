package main

import "fmt"

// fallthrough inseamna ca desi indeplineste conditia respectiva, va RULA SI URMATOAREA CHIAR DACA SE POTRIVESTE CU TIPARUL.
func main() {
	switch "daniel" {
	case "daniel":
		fmt.Println("warap daniel")
		fallthrough
	case "gigi":
		fmt.Println("PRINTEAZA SI ASTA")
	default:
		fmt.Println("Have you no fr?")
	}
}
