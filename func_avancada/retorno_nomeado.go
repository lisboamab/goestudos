package main

import "fmt"

func somaESubtrai(n1, n2 int) (soma int, sub int) {
	soma = n1 + n2
	sub = n1 - n2
	return
}

func main() {
	fmt.Println(somaESubtrai(10, 20))
	soma, subtracao := somaESubtrai(10, 20)
	fmt.Println(soma, subtracao)
}