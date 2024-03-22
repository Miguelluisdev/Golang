package main

import "fmt"

func main() {
	test := "test"

	switch test {
	case "teste":
		fmt.Print("opção1")
		fallthrough

	case "teste2":
		fmt.Print("opção2")

	default:
		fmt.Print("oplçao3")
	}

}
