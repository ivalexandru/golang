// the final parameter in a function signature may have a TYPE prefixed with "..."
// a function with such a parameter is called variadic, and MAY BE INVOKED WITH ZERO OR MORE ARGUMENTS FOR THAT PARAMETER.

package main

import "fmt"

func main() {
	n := average(43, 56, 87, 12, 45, 57)
	fmt.Println(n) //50
	fmt.Println("\n")
	//////////////////////////////////////////
	xxx := media(2, 4) //3
	fmt.Println(xxx)
}

func average(sf ...float64) float64 {
	fmt.Println(sf)
	fmt.Printf("%T", sf)
	var total float64 // sets total to the zero value
	for _, v := range sf {
		total += v
	}
	return total / float64(len(sf))
}

// The range form of the for loop iterates over a slice or map.
// When ranging over a slice, two values are returned for each iteration. The first is the index, and the second is a copy of the element at that index.
// RANGE on arrays and slices provides both the index and value for each entry. Above we didn’t need the index, so we ignored it with the blank identifier _.
// Sometimes we actually want the indexes though.(si atunci putem pune i sau ceva in loc de _ )

func media(valoare ...float64) float64 {
	var totalul float64
	for _, val := range valoare {
		totalul += val
	}
	return totalul / float64(len(valoare))
}
