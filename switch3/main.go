package main

import "fmt"

//ASTA NU ARE CONDITIE DE VERIFICAT, ruleaza PRIMUL case care este adevarat.
var variabila string = "pa"

func main() {
	switch {
	case len(variabila) == 2:
		fmt.Println("PRINTEAZA ASTA")
	case variabila == "gigi":
		fmt.Println("nu printeaza")
	default:
		fmt.Println("nimic nu e adevarat, so print this")
	}
}
