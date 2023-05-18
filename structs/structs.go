package main

import "fmt"

type user struct {
	nome string
	idade uint8
	email string
	endereco endereco
}

type endereco struct {
	logradouro string
	numero int
}

func main() {
	//Structs em go são similares as classes nas linguagens OO
	fmt.Println("Arquivo Structs")

	var u user
	u.nome = "Roberto"
	u.idade = 21
	u.email = "roberto@tuamaeaquelaursa.com.br"
	fmt.Println(u)

	endereco2 := endereco{"Rua da Quitanda", 2}
	usuario2 := user{"Davi", 22, "davi@tuamaeaquelaursa.com.br", endereco2}
	fmt.Println(usuario2)

	endereco3 := endereco{logradouro: "Rua das Amoras", numero: 33}
	usuario3 := user{email: "semnome@tuamaeaquelaursa.com.br", idade: 23, endereco: endereco3}
	fmt.Println(usuario3)
	fmt.Printf("E-mail: %v \nIdade: %v \nRua: %v, %v\n", usuario3.email, usuario3.idade, usuario3.endereco.logradouro, usuario3.endereco.numero)
}