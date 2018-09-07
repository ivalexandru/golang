package main

import "fmt"

func main() {
	mySlice := make([]int, 1)
	fmt.Println(mySlice[0]) //0
	// fmt.Println(mySlice[3]) // va da eroare, pt ca am length 1, daca vreau sa adaug pana la 3, tre sa dau cu append
	mySlice[0] = 7
	fmt.Println(mySlice[0]) //7
	//daca vreau sa adaug 1 la valoarea existenta la indexul 0 al sliceului:
	mySlice[0]++
	fmt.Println(mySlice[0]) // 8

}
