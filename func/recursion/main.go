package main

import "fmt"

func factorial(x int) int {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1) //ASTA TOT SCADE CATE 1, PANA AJUGE LA  X = 0, CAZ IN CARE IF-UL DE SUS DA RETURN SI NU SE MAI AJUNGE LA ACEST RETURN.
	// 4 nuy e zero, ajungem la al doilea return, 4 * 3, apoi de la incaput factorial(3), ajungem iarasi jos (DAR DE DATA ASTA NU MAI INMULTIM CU 4, INTRAM DIRECT DUPA INMULTIRE SI 3-1 => 2, DECI AVEM  4 * 3 * 2 SI TOT ASA PANA CAND X == 0)
}

func main() {
	fmt.Println(factorial(4))
}
