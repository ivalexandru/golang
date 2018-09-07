package main

import "fmt"

func main() {

	mapaMea := map[string]bool{
		"salome": true,
		"marian": true,
		"gigi":   true,
	}

	if mapaMea["salome"] {
		fmt.Println("salome e in mapa")
	}

	fmt.Println("toata mapa este asta: ", mapaMea)

}
