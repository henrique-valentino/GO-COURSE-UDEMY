package main

import (
	"errors"
	"fmt"
)

func main() {
	//numeros inteiros
	//int8, int16, int32, int64
	var numero int16 = 100
	fmt.Println(numero)

	// uint sem sinal

	//numeros reais
	// float32, float64
	var numeroreal float32 = 1.32
	fmt.Println(numeroreal)

	// string
	var str string = "texto"
	fmt.Println(str)

	char := 'B'
	fmt.Println(char)

	// valor inicial para uma variavel
	// string -> string vazia
	// numero -> 0
	// tipo de dado -> nill
	// boolean -> false

	var booleano bool = true
	fmt.Println(booleano)

	var erro error = errors.New("Erro interno")
	fmt.Println(erro)

}
