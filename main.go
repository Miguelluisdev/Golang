package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	fn, err := printFullName("darius")
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(fn)
}

func printFullName(p string) (string, error) {
	if p == "" {
		return "", errors.New("o nome não pode estar vazio")
	}
	return p, nil
}
