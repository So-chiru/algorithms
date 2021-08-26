package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var input = ""
	fmt.Scanln(&input)

	number, err := strconv.Atoi(input)
	if err != nil {
		os.Exit(1)
	}

	if number%2 == 0 && number != 2 {
		fmt.Println("YES")
		return
	}

	fmt.Println("NO")
}
