package main

import "fmt"

func main() {
	retorno := func(texto string, numero int) string {
		return fmt.Sprintf("Recebido -> %s\t%v", texto, numero)
	}("Passando Parâmetro", 10)

	fmt.Println(retorno)
}