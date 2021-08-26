package main

import (
	"fmt"
)

func main() {
	var length int
	fmt.Scanf("%d\n", &length)

	for {
		var input = ""
		fmt.Scanln(&input)

		if len(input) == 0 {
			break
		}

		if len(input) <= 10 {
			fmt.Println(input)
			continue
		}

		fmt.Println(string(input[0]) + fmt.Sprint(len(input)-2) + string(input[len(input)-1]))
	}
}
