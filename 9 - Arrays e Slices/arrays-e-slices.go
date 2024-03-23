package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Arrays e Slices")

	var array1 [5]string
	array1[0] = "Posição 1"
	fmt.Println(array1)

	array2 := [2]string{"posição 1", "posição 2"}
	fmt.Println(array2)

	array3 := [...]string{"posição 1", "posição 2"}
	fmt.Println(array3)

	slice1 := []int{1, 2, 3}
	fmt.Println(slice1)

	fmt.Println(reflect.TypeOf(slice1))
	fmt.Println(reflect.TypeOf(array1))

	slice1 = append(slice1, 4)
	fmt.Println(slice1)

	arrayInt := [5]int{1, 2, 3, 4, 5}
	slice2 := arrayInt[1:3]
	fmt.Println(slice2)

	// arrays internos
	slice3 := make([]float32, 10, 15)
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

	slice4 := make([]float32, 5)
	fmt.Println(slice4)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))

}
