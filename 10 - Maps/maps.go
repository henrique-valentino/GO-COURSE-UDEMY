package main

import "fmt"

func main() {
	fmt.Println("Maps")
	// [type key] type values
	usuario := map[string]string{
		"nome":      "Henrique",
		"sobrenome": "Valentino",
	}
	fmt.Println(usuario)
	fmt.Println(usuario["nome"])

	delete(usuario, "sobrenome")
	fmt.Println(usuario)

	usuario["sobrenome"] = "Valentino"
	fmt.Println(usuario)

}
