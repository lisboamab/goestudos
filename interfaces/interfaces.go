package main

import (
	"fmt"
	"math"
)

type forma interface {
	area() float64
}

//caso seja passado uma interface vazia, essa função poderá ser passada com qualquer interface
func escreverArea(self forma) {
	fmt.Printf("A area da forma é %0.2f\n", self.area())
}

type retangulo struct {
	altura float64
	largura float64
}

func (self retangulo) area() float64 {
	return self.altura * self.largura
}

type circulo struct {
	raio float64
}

func (self circulo) area() float64{
	return math.Pi * math.Pow(self.raio, 2)
}

func main() {
	circulo1 := circulo{10}
	escreverArea(circulo1)

	retangulo1 := retangulo{10, 15}
	escreverArea(retangulo1)
}