package main

import (
	"fmt"
)

func main() {
	var problems int
	fmt.Scanf("%d\n", &problems)

	var solvable int = 0

	for i := 0; i < problems; i++ {
		var petya int
		var vasya int
		var tonya int

		fmt.Scanf("%d %d %d\n", &petya, &vasya, &tonya)

		if petya+vasya+tonya > 1 {
			solvable++
		}
	}

	fmt.Println(solvable)
}
