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
	delete(myGreeting, 2)
	fmt.Println(myGreeting)

}
