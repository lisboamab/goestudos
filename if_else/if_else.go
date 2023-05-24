package main

import "fmt"

func main(){
	fmt.Println("Estruturas de controle")

	numero := 10

	if numero > 115 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor que 15")
	}

	//no caso abaixo a variavel outroNumero se limita ao escopo da estrutura do controle é indicado para variaveis volateis
	// não existe if ternario no go a escrita abaixo representa um 'if statement initialization'
	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("Número é maior que zero")
	} else if numero < -10{
		fmt.Println("Numero é menor que -10")
	} else {
		fmt.Println("Entre -10 e 0")
	}
}