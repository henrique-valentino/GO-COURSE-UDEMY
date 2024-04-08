package main

import "fmt"

func generica(inter interface{}) {
	fmt.Println(inter)
}

func main() {
	generica("String")
	generica(1)
	generica(true)
}
