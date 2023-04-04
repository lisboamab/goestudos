package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main(){
	fmt.Println("Escreva do arquivo main")
	auxiliar.Escrever()

	erro := checkmail.ValidateFormat("123")

	fmt.Println(erro)
}