package main

import "fmt"

var userInput string

func main() {
	fmt.Println("enter your name:  ")
	fmt.Scan(&userInput)
	fmt.Println("te cheama ", userInput)
	fmt.Println("te cheama AND", &userInput)
}

//MINUTUL 5555555555555555555
