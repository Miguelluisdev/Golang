package main

func somaVariosNumeros(numeros ...int) int {
	total := 0
	for _, num := range numeros {
		total += num
	}
	return total
}

func main() {

}
