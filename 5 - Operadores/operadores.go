package main

import "fmt"

func main() {
	//Aritimeticos
	// + - * / %

	var numero1 int32 = 10
	var numero2 int32 = 15
	var resultado int32 = numero1 + numero2
	fmt.Println(resultado)

	// Atrinuição
	var variavel1 string = "String"
	variavel2 := "String2"
	fmt.Println(variavel1, variavel2)

	// operadores relacionais
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 < 2)

	//Logicos
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)
	fmt.Println(!falso)

	//Unarios
	// ++ --
	numero := 10
	numero++
	numero += 15
	fmt.Println(numero)

	//não tem ternario

}
