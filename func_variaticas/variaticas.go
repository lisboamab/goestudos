package main

import "fmt"

func soma(numeros ...int) int { // isso permite passar vários parametros daquele tipo na função
	var total int
	for _, numero := range numeros {
		total += numero
	}
	return total
}

func escrever(texto string, numeros ...int) { //só pode ter um variatico por função
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}

func main() {
	totalDaSoma := soma(1, 2, 3, 5, 8)
	fmt.Println(totalDaSoma)
	escrever("Teste", 1, 2, 3, 4, 5)
}