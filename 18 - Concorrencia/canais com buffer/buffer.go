package main

import "fmt"

func main() {
	canal := make(chan string, 2)
	canal <- "Olá mundo!"
	canal <- "Progamando em GO"

	mensagem := <-canal
	fmt.Println(mensagem)
}
