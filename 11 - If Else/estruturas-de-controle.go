package main

import "fmt"

func main() {
	fmt.Println("Estruturas de controle")

	numero := 10

	if numero > 15 {
		fmt.Println("maior q 15")
	} else {
		fmt.Println("Menor ou igual a 15")
	}

	// so vale dentro do escopo do if/else
	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("numero criado dentro do if")
	}
}
