package main

import (
	"errors"
	"fmt"
)

func main() {

	//tipos inteiros
	var numero1 int8 = 100
	fmt.Println(numero1)

	var numero2 int16 = 100
	fmt.Println(numero2)

	var numero3 int32 = 100
	fmt.Println(numero3)

	var numero4 int64 = 100
	fmt.Println(numero4)

	//numeros decimais
	var decimal1 float32 = 10.12
	fmt.Println(decimal1)

	var decimal2 float64 = 9797989.999
	fmt.Println(decimal2)

	decimal3 := 89768.887
	fmt.Println(decimal3)

	//Texto

	var texto string = "Meu texto"
	fmt.Println(texto)

	//Outros tipos

	var boleano bool
	fmt.Println(boleano)

	var boleano1 bool = true
	fmt.Println(boleano1)

	var erro error = errors.New("Erro interno")
	fmt.Println(erro)
}
