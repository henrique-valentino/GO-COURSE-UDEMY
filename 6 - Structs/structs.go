package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
}

func main() {
	fmt.Println("Arquivo Structs")

	var u usuario
	u.nome = "Davi"
	u.idade = 21
	fmt.Println(u)
	endereco := endereco{"Rua"}
	u2 := usuario{"Hnerique", 42, endereco}
	fmt.Println(u2)

	u3 := usuario{nome: "Henrique"}
	fmt.Println(u3)

	u4 := usuario{idade: 15}
	fmt.Println(u4)

}
