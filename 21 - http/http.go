package main

import (
	"log"
	"net/http"
)

func usuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Pagina Usuários"))
}

func main() {
	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Olá Mundo!"))
	})

	http.HandleFunc("/usuarios", usuarios)

	// sobe servidor
	log.Fatal(http.ListenAndServe(":5000", nil))
}
