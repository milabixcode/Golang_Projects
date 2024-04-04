package main

import (
	"fmt"
	"math"
)

type forma interface {
	area() float64
}

func escreverArea(f forma) {
	fmt.Printf("A área da forma é %0.2f\n", f.area())
}

type retangulo struct {
	altura  float64
	largura float64
}

type circulo struct {
	raio float64
}

func (r retangulo) area() float64 {
	return r.altura * r.largura
}

func (c circulo) area() float64 {
	return math.Pi * c.raio*c.raio
}

func main() {
	retangulo1 := retangulo{10,20}
	escreverArea(retangulo1)

	circulo1 := circulo{2}
	escreverArea((circulo1))



}
