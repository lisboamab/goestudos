package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Arrays and Slices")

	var array1 [5]string // declara a variavel de array, especifica o tamanho e o tipo de dado
	array1[0] = "posição 1"
	fmt.Println(array1)

	array2 := [5]string{"posição1", "posição2", "posição3", "posição4", "posição5"}

	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5} //arrays em go não são dinamicos, o número de elementos é fixo.
	fmt.Println(array3)

	slice := []int{1, 2, 56, 77, 99} //slices são dinamicos, é mais usado do que arrays em go.
	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array3))

	slice = append(slice, 18) 
	fmt.Println(slice)

	slice2 := array2[1:3] //o primeiro indice é inclusivo e o segundo indice é exclusivo

	array2[1] = "posição alterada"
	fmt.Println(slice2)

	//Arrays Internos
	slice3 := make([]float32, 10, 11) //a função make cria um array com 15 posições nesse caso, e pega as 10 primeiras posições
	slice3 = append(slice3, 5)
	slice3 = append(slice3, 6)
	
	
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))


}