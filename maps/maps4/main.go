package main

import "fmt"

func main() {

	myGreeting := map[int]string{
		0: "Good m",
		1: "bon jur",	
		2: "buen d",
		3: "bgiorno",
	}

	fmt.Println(myGreeting)
	// delete(myGreeting, 2)
	fmt.Println(myGreeting)

	// delete(myGreeting, 2)

	if val, exista := myGreeting[2]; exista {

		fmt.Println("val: ", val)
		fmt.Println("exists ", exista)
	} else {
		fmt.Println("that value does not exist")
		fmt.Println("val: ", val)
		fmt.Println("exista ", exista)
	}
}
