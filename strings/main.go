package main

import "fmt"

func salam(par int) int {
	return par * 2
}

// A STRING IS A COLLECTION OF RUNES
//cand folosesti BACKTICKS pt string, pastreaza returns/enter si formatarea, cand folosesti GHILIMELE nu le pastreaza.
func main() {
	fmt.Println(salam(2))
}

//daca pui ghilimele mai sus in loc de backticks iei si eroare pt ca ai spatiu...
