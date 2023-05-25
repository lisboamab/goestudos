package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0

	// primeira forma de escrever um for
	for i < 10 {
		i++
		fmt.Println(i)
		time.Sleep(time.Second)
	}

	fmt.Println(i)

	// segunda forma de fazer um for
	for j := 0; j < 10; j++ {
		fmt.Printf("Incrementando j: %v\n", j)
		time.Sleep(time.Second)
	}

	nomes := [3]string{"João", "Davi", "Lucas"}
	nomes2 := []string{"Zé", "João", "Nestor"}
	
	//  posição, elemento
	for indice, nome := range nomes {
		fmt.Printf("%v: %v\n", indice, nome)
	}

	for i, nome := range nomes2{
		fmt.Printf("%v:\t%v\n", i, nome)
	}
}