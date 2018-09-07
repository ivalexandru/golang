package main

import "fmt"

func main() {

	strArr := []string{"a", "b", "c", "d", "e"}
	for index, caracter := range strArr {
		fmt.Print("indexul:  ", index)
		fmt.Println("  caracterul:  ", caracter)
	}

	fmt.Println(strArr)
	strArr = append(strArr[:1], strArr[3:]...)
	fmt.Println(strArr) // au disparut indexul 1 si indexul 2

}
