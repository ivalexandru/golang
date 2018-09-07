package main

import (
	"fmt"
)

//TZ E ORICE KEY
//IF TZ IZ PREZENT, SECONDS WILL BE SET APPROPRIATELY AND OK WILL BE TRUE. IF NOT, SECONDS WILL BE SET TO ZERO AND OK WILL BE FALSE

var timeZone = map[string]int{
	"UTC": 20,
	"EST": 30,
}

func offset(tz string) int {
	if seconds, ok := timeZone[tz]; ok {
		fmt.Println("ok: ", ok) //true
		return seconds
	}
	fmt.Println("unknown time zone: ", tz)
	return 0
}

func main() {

	fmt.Println(offset("UTC"))
	// seconds:  20
	// ok:  true
	// 20

	fmt.Println("=======================")

	fmt.Println(offset("salam"))
	//unknown time zone:  salam
	// 0

}
