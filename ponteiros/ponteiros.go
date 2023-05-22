package main

import "fmt"

func main() {
	fmt.Println("Ponteiros")

	var variavel1 int = 10
	var variavel2 int = variavel1
	variavel3 := variavel2

	fmt.Println(variavel1, variavel2, variavel3)
	variavel1++
	fmt.Println(variavel1, variavel2, variavel3)
	variavel2 = variavel2 + 2
	fmt.Println(variavel1, variavel2, variavel3)

	var variavel4 int
	var ponteiro *int //essa variavel vai receber uma referencia

	//COM PONTEIROS
	variavel4 = 100
	ponteiro = &variavel4

	fmt.Println(variavel4, ponteiro)
	fmt.Println(variavel4, *ponteiro)

}