package main

import "fmt"

type pessoa struct {
	nome   string
	idade  uint8
	altura uint8
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	fmt.Println("Herança")
	p1 := pessoa{"Henrique", 42, 170}

	e2 := estudante{p1, "BSI", "UCB"}
	fmt.Println(e2)
	fmt.Println(e2.nome)

}
