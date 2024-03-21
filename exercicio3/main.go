package main

import (
	"fmt"
)

func somaSlice(slice []int) int {
	soma := 0

	for _, num := range slice {
		soma += num
	}
	return soma
}

func main() {
	numeros := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("a soma dos numeros no slice é :", somaSlice(numeros))
}
