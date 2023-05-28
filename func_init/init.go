package main

import "fmt"
var n int
//ela é executada sempre antes da função main()
//pode ter uma por arquivo
//talvez seja bom para setar variaveis globais
func init(){
	fmt.Println("Executando a função Init")
	n = 10
}

func main() {
	fmt.Println(n)
	fmt.Println("Função Init")
}