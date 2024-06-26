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
		return "Sabado-Feira"
	default:
		return "Data invalida"
	}
}

func fimDeSemana(numero int) string {
	switch {
	case numero == 1 || numero == 7:
		return "Fim de semana"
	default:
		return "Dia de semana"
	}
}

func main() {
	fmt.Println("Switch")
	dia := diaDaSemana(3)
	fmt.Println(dia)

	fds := fimDeSemana(2)
	fmt.Println(fds)
}
