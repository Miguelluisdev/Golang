package main

import (
	"fmt"
)

func main() {
	var numbers [10]float64
	var sum float64

	for i := 0; i < 10; i++ {
		fmt.Printf("Digite um numero %d : ", i+1)
		fmt.Scanln(&numbers[i])
		sum += numbers[i]
	}

	average := sum / 2

	fmt.Printf("a conta é %2.f", average)

}
