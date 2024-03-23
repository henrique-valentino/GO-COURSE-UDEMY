package main

import "fmt"

func alunoEstaAprovado(n1, n2 float32) bool {
	defer fmt.Println("Media calculada. Resultados será retornado!")
	fmt.Println("Executadndo função alunoEstaAprovado")
	media := (n1 + n2) / 2

	if media >= 6 {
		return true
	}

	return false
}

func main() {
	fmt.Println(alunoEstaAprovado(7, 8))
}
