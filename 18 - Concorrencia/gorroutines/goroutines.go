package main

import (
	"fmt"
	"time"
)

func main() {
	// concorrencia != paralelismo
	go escrever("Olá Mundo!") // goroutines
	escrever(("Programando em GO!"))
}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
