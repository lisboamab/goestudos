package main

import "fmt"

func main() {
	fmt.Println("Maps")
// dentro dos colchetes é o tipo das keys e fora dos values
// maps são implementações de hash pro go
	usuario := map[string]string {
		"nome": 			"Pedro",
		"sobrenome": "Lisboa",
	}

	fmt.Println(usuario["nome"])

	//map aninhado
	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "Marcos",
			"segundo": "Antonio",
		},
		"curso": {
			"nome": "Banco de Dados",
			"campus": "Maraponga",
		},
	}

	fmt.Println(usuario2)
	fmt.Println(usuario2["nome"]["primeiro"])

	delete(usuario2, "nome")

	usuario2["signo"] = map[string]string {
		"signo": "aries",
	}

	fmt.Println(usuario2)
}