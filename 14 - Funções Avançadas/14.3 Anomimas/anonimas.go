package main

import "fmt"

func main() {

	retorno := func(texto string) string {
		return fmt.Sprintf("recebendo texto-> %s", texto)
	}("Passando Parametro")

	fmt.Println(retorno)

}
