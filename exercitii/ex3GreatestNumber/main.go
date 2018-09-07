package main

import "fmt"

func main() {
	greatest := findGreatest(1, 2, 3)
	fmt.Println(greatest)

}

func findGreatest(lista ...int) int {
	var largest int
	for _, i := range lista {
		if i > largest {
			largest = i
		}
	}
	return largest
}
