package main

import "fmt"

func soma(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}

func main() {
	result := soma(1, 2, 3, 4, 5)
	fmt.Println(result)

}
