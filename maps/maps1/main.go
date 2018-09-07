package main

import "fmt"

MAPS - AN UNORDERED group of elements of one type, indexed by a set of key.
the value of an uninitialized map is nil
a map is built on top of a hashtable
map is a refference type.

func main() {
	//pt a creea EMPTY map:
	mapaGoala := make(map[int]string)
	fmt.Println(mapaGoala) //map[]
	mapaGoala[1] = "salam"
	fmt.Println(mapaGoala) // map[1:salam]

	//to create and initialize a new map:
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println(n) // map[foo:1 bar:2]
}
