package main

import (
	"fmt"
	"math"
)

type Forma interface {
	Area() float64
}

type Retangulo struct {
	Base   float64
	Altura float64
}

type Circulo struct {
	Raio float64
}

func (r Retangulo) Area() float64 {
	return r.Base * r.Altura
}

func (c Circulo) Area() float64 {
	return math.Pi * c.Raio * c.Raio
}

func main() {
	formas := []Forma{
		Retangulo{Base: 5, Altura: 10},
		Circulo{Raio: 4},
	}
	for _, formas := range formas {
		area := formas.Area()
		fmt.Printf("Area das formas %2f\n", area)
	}
}
