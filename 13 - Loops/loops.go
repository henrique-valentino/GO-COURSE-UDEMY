package main

import (
	"fmt"
)

func main() {
	i := 0

	for i < 10 {
		i++
		fmt.Println("Incrementando i")
	}
	fmt.Println(i)

	for j := 0; j < 10; j++ {
		fmt.Println("Incrementando j", j)
	}

	for j := 0; j < 10; j += 2 {
		fmt.Println("Incrementando j", j)
	}

	nomes := [3]string{"João", "Paulo", "Henrique"}

	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}

	for _, nome := range nomes {
		fmt.Println(nome)
	}

	usuario := map[string]string{
		"nome":      "Henrique",
		"sobrenome": "Valentino",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

	for chave := range usuario {
		fmt.Println(chave)
	}

	for _, valor := range usuario {
		fmt.Println(valor)
	}

}
