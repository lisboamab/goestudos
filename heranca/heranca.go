package main

import "fmt"

type people struct{
	nome string
	idade uint8
	email string
}

type student struct{
	//aplica o conceito de "herança" no go, embora não seja uma linguagem OO
	people
	matricula string

}

func main(){
	fmt.Println("Herança")
	
	var pessoa people
	pessoa.nome = "Marcos"
	pessoa.idade = 26
	pessoa.email = "marcos@tuamaeaquelaursa.com.br"

	pessoa1 := people{nome: "Xuxa", idade: 56, email: "xuxa@show.se"}


	fmt.Printf("Pessoa : %v\n", pessoa)
	fmt.Printf("Pessoa 1: %v\n", pessoa1)

	estudante := student{people: pessoa, matricula: "T345444"}

	fmt.Printf("%v",estudante)
	
}