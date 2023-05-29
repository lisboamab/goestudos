package main

import "fmt"

type usuario struct {
	nome string
	idade uint8
}

/*
- o 'u' seria algo como um self
- a sintaxe a baixo é a de um metodo
*/
func (u usuario) salvar() {
	fmt.Printf("Salvado os dados do Usuario %s de idade %d\n", u.nome, u.idade)
}

func (self usuario) maiorDeIdade() bool {
	return self.idade >= 18
}

func (self *usuario) fazerAniversario() {
	self.idade++
}

func escrever() {
}

func main () {
	usuario1 := usuario{"Marcos", 26}
	usuario2 := usuario{nome: "João", idade: 30}
	usuario1.salvar()
	usuario2.maiorDeIdade()
	
	fmt.Println(usuario2.idade)
	usuario2.fazerAniversario()
	fmt.Println(usuario2.idade)
}