package main 

import "fmt"

func recuperarExecucao() {
	if r := recover(); r != nil {
		fmt.Println("Execução recuperada")
	}
}
func alunoEstaAprovado(n1, n2 float64) bool {
	defer recuperarExecucao()
	media := (n1 + n2)/2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}
	
	// é uma forma de tratamento de erro porém, não é a melhor
	panic("MEDIA EXATAMENTE 6")
}

func main() {
	fmt.Println(alunoEstaAprovado(6, 8))
	fmt.Println("Só vou executar quando o panic for tratado")
}