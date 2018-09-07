package main

import (
	"fmt"
)

func main() {
	mySlice := make([]int, 0, 5)
	fmt.Println(mySlice)      // pt ca len mai sus e zero, printeaza un slice gol  []
	fmt.Println(len(mySlice)) // 0
	fmt.Println(cap(mySlice)) // 5
	mySlice = append(mySlice, 9)

	myOtherSlice := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(myOtherSlice)

	//APENDUIESC UN SLICE LA ALT SLICE, folosind TREI PUNCTE:
	myOtherSlice = append(mySlice, myOtherSlice...)
	fmt.Println(myOtherSlice) // [9 1 2 3 4 5 6]

	//PT A STERGE FOLOSESTI TOT APPEND
	//DACA VREI SA STERGI DE LA INCEPUT PANA LA POZITIA 1(exclusiv), apoi de la pozitia 2 pana la final:
	myOtherSlice = append(myOtherSlice[:1], myOtherSlice[3:]...)
	fmt.Println(myOtherSlice) // [9 3 4 5 6]

}

// Len: 1 Capacity: 5 Value: 0
// Len: 2 Capacity: 5 Value: 1
// Len: 3 Capacity: 5 Value: 2
// Len: 4 Capacity: 5 Value: 3
// Len: 5 Capacity: 5 Value: 4
// Len: 6 Capacity: 10 Value: 5
// Len: 7 Capacity: 10 Value: 6
// Len: 8 Capacity: 10 Value: 7
// Len: 9 Capacity: 10 Value: 8
// Len: 10 Capacity: 10 Value: 9
// Len: 11 Capacity: 20 Value: 10
