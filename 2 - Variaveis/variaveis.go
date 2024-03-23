package main

import "fmt"

func main() {
	var variavel string = "Variavel 1"
	variavel2 := "Variavel 1"

	fmt.Println(variavel)
	fmt.Println(variavel2)

	var (
		variavel3 string = "String 1"
		variavel4 string = "String 2"
	)

	fmt.Println(variavel3, variavel4)

	variavel5, variavel6 := "String 5", "String 6"

	fmt.Println(variavel5, variavel6)

	const constante string = "Constante"

	fmt.Println(constante)
}
