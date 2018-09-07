package main

import "fmt"

func main() {
	switch "Mehdi" {
	case "daniel":
		fmt.Println("warap daniel")
	case "Mehdi":
		fmt.Println("warap Mehdi")
	default:
		fmt.Println("Have you no fr?")
	}
}
