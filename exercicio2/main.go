package main

import (
	"fmt"
)

func main() {
	var number [5]int
	max := 0

	for i := 0; i < 5; i++ {
		fmt.Printf("Digite um numero %d", i+1)
		fmt.Scanln(&number[i])

		if number[i] > max {
			max = number[i]
		}

	}

	fmt.Printf("numero maior é %d", max)
}
