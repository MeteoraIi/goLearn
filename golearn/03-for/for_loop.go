package main

import (
	"fmt"
)

func main() {
	i := 1

	for {
		fmt.Println("loop")
		break
	}

	for j := 1; j < 9; j++ {
		fmt.Print(j, " ")
	}

	fmt.Println()

	for n := 0; n < 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Print(n, " ")
	}

	fmt.Println()

	for i < 3 {
		fmt.Print(i, " ")
		i++
	}
}
