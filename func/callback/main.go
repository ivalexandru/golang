package main

import "fmt"

//  []int = a slice of int
// func(int)  =a function that takes an int as arg

func visit(numbers []int, callback func(int)) {

	//ia toate numerele din slice-ul primit si ruleaza functia callback pe ele:
	for _, n := range numbers {
		callback(n)
	}
}

func main() {
	visit([]int{1, 2, 3, 4}, func(n int) {
		// fmt.Println(n) //aici pui ce vrei sa faca functia "callback"
		fmt.Println(n * 2)
	})
}

//callback nu e keywork, poti pune ce nume doresti.
