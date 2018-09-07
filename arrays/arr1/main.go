package main

import "fmt"

// Array = numbered sequence of elements OF A SINGLE TYPE.
// Does NOT CHANGE SIZE => nu poti adauga chestii(nu sunt dinamice). nu prea le folosesti, le folosesc slice-urile.

// var [58] array - daca are cifre intre [ ] este array, altfel e slice.\
// si [58] face parte din TYPE, impreuna cu array
// 58 e LENGTH

func main() {
	var x [58]string
	fmt.Println(x)
	// [                                                         ]
	//estimez ca 58 de spatii sunt acolo
	fmt.Println(len(x)) //58
	fmt.Println(x[33])  //itemul de la pozitia 33
}
