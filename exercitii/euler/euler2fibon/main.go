// Each new term in the Fibonacci sequence is generated by adding the previous two terms. By starting with 1 and 2, the first 10 terms will be:

// 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...

// By considering the terms in the Fibonacci sequence whose values do not exceed four million, find the sum of the even-valued terms.

package main

import (
	"fmt"
)

func main() {

	fmt.Println(fibonacci(6)) // 8

	fmt.Println(FibonacciLoop(3)) // 2  ( 0 + 1 + 1)

}

//cu slice si loops:

func FibonacciLoop(n int) int {
	f := make([]int, n+1, n+2)
	if n < 2 {
		f = f[0:2]
	}
	f[0] = 0
	f[1] = 1
	for i := 2; i <= n; i++ {
		f[i] = f[i-1] + f[i-2]
	}
	return f[n]
}

// We are working with slices because we cannot create arrays with numeric lengths that are not constant. The first thing we do is create the slice which has a length of n+1 and a capacity of n+2. The length needs to be more than n because if we pass in a value of two, we also need to accommodate the zero index which means our slice needs to be of length three, not two.

// If our placeholder is less than two, we need to increase the length of our slice to hold at least the first and second Fibonacci numbers. Finally we can loop similarly to how we did with recursion and return the final value of the slice. Each resulting value in the loop is stored in the slice.

// cu recursion:
func fibonacci(x int) int {
	if x <= 1 {
		return x
	}
	return fibonacci(x-1) + fibonacci(x-2)
}