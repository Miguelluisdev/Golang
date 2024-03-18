package main

import "fmt"

func Arauto() {
	meuMapa := make(map[string]bool)

	meuMapa["darius"] = true
	meuMapa["cidk"] = true
	meuMapa["aura"] = false

	fmt.Printf("primitivo=%+v\n", meuMapa)
}
