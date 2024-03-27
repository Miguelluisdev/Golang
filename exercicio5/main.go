package main

import (
	"fmt"
	"math/rand"
	"time"
)

// func oneExecult() {
// 	time.Sleep(1000 * time.Millisecond)
// }

func advinhaNumero() {
	var numero = 60
	var tentativas = 0
	for tentativas < 21 {
		var tentativa = rand.Intn(99) + 1
		tentativas++
		if tentativa < numero {
			fmt.Println("o numero está a baixo", tentativa)
		} else if tentativa > numero {
			fmt.Println("o numero está acima", tentativa)
		} else {
			fmt.Println(" numero ceto miseravel", tentativa)
			break
		}

	}
	if tentativas > 20 {
		fmt.Print("numero limite de tentivas")
	}
	time.Sleep(1000 * time.Minute)
}

func main() {
	advinhaNumero()
}
