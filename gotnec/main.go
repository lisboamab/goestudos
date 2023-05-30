package main

import (
	"fmt"
	"gotnec/app"
	"os"
	"log"
)

func main() {
	fmt.Println("Ponto de partida")

	aplicacao := app.Gerar()
	if erro := aplicacao.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}
}