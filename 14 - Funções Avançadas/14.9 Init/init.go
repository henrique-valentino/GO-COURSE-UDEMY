package main

import "fmt"

var n int

func init() {
	fmt.Println("Função init")
	n = 10
}

func main() {
	fmt.Println("Funçã main")
	fmt.Println(n)
}
