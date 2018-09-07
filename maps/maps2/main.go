package main

import "fmt"

func main() {
	// [string] reprezinta key
	// al doilea string reprezinta type-ul pt valorile care este
	var myGreeting = make(map[string]string)
	myGreeting["Tim"] = "Good M"
	myGreeting["Jenny"] = "bonjr"
	fmt.Println(myGreeting)

	// COMPOSITE LITERAL (literal as in literal values - COMPOSED OF VALUES)
	//NU UITA SA PUI DOUA PUNCTE
	salam := map[string]string{
		"gigi":  "neata!",
		"Jenny": "buna",
	}
	salam["Hara"] = "bau" //adaugi intrari
	fmt.Println(salam)
}
