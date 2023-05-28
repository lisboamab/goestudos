package main

import (
	"fmt"
)

func funcao1() {
	fmt.Println("Executando função 1")
}

func funcao2() {
	fmt.Println("Executando função 2")
}

func alunoAprovado(n1, n2 float32) bool {
	var media float32 = (n1 + n2)/2 
	fmt.Println("Entrando na função de aluno aprovado")
	defer fmt.Printf("Media: %v\n", media)
	return media > 5
}

func main() {
	// o defer 'adia' a execução de determinada chamada, até o ultimo momento possivel, no caso de uma função, até o retorno da mesma
	defer funcao1()
	funcao2()

	fmt.Println(alunoAprovado(4, 7))
}