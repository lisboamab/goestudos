package main

import "fmt"

func inverteSinal(numero int) int {
	return numero * -1
}

func inverteSinalComPonteiro(numero *int){
	/*
	- funções com ponteiro sobrescrevem o endereço de memoria
	- não possuem retorno
	- usam essa sintaxe de parametro sendo passado com & e asterisco nas chamadas dos mesmos
	*/
	*numero = *numero * -1
}

func main() {
	numero := 3

	fmt.Println(numero)
	numeroInvertido := inverteSinal(numero) //a variavel não foi alteradag
	fmt.Println(numeroInvertido)
	fmt.Println(numero)

	novoNumero := 40
	fmt.Println(novoNumero)
	inverteSinalComPonteiro(&novoNumero) //a variavel foi alterada
	fmt.Println(novoNumero)
}