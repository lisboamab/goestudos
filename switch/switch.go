package main

import "fmt"

func diaDaSemana(numero int) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda-Feira"
	case 3:
		return "Terça-Feira"
	case 4:
		return "Quarta-Feira"
	case 5:
		return "Quinta-Feira"
	case 6:
		return "Sexta-Feira"
	case 7:
		return "Sabádo"
	default:
		return "ERROR: Numero Invalido"
	}
}

func diaDaSemana2(numero int) string {
	switch {
	case numero == 1:
		return "Domingo"
	case numero == 2:
		return "Segunda-Feira"
	case numero == 3:
		return "Terça-Feira"
	case numero == 4:
		return "Quarta-Feira"
	case numero == 5:
		return "Quinta-Feira"
	case numero == 6:
		return "Sexta-Feira"
	case numero == 7:
		return "Sabádo"
	default:
		return "ERROR: Numero Invalido"
	}
}

func diaDaSemana3(numero int) string {
	var diaEmTexto string
	
	switch {
	case numero == 1:
		diaEmTexto = "Domingo"
		// fallthrough executa a proxima condição
	case numero == 2:
		diaEmTexto = "Segunda-Feira"
	case numero == 3:
		diaEmTexto = "Terça-Feira"
	case numero == 4:
		diaEmTexto = "Quarta-Feira"
	case numero == 5:
		diaEmTexto = "Quinta-Feira"
	case numero == 6:
		diaEmTexto = "Sexta-Feira"
	case numero == 7:
		diaEmTexto = "Sabádo"
	default:
		diaEmTexto = "ERROR: Numero Invalido"
	}

	return diaEmTexto
}

func main()	{
	fmt.Println(diaDaSemana(1))
	fmt.Println(diaDaSemana2(1))
	fmt.Println(diaDaSemana3(1))

	dia := diaDaSemana(10)
	fmt.Println(dia)

}