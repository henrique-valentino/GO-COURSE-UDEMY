package main

import (
	"crud/servidor"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const PATH_USUARIO = "/usuarios"
const PATH_USUARIO_ID = "/usuarios/{id}"

func main() {
	router := mux.NewRouter()
	router.HandleFunc(PATH_USUARIO, servidor.CriarUsuario).Methods(http.MethodPost)
	router.HandleFunc(PATH_USUARIO, servidor.BuscarUsuarios).Methods(http.MethodGet)
	router.HandleFunc(PATH_USUARIO_ID, servidor.BuscarUsuario).Methods(http.MethodGet)
	router.HandleFunc(PATH_USUARIO_ID, servidor.AtualizarUsuario).Methods(http.MethodPut)
	router.HandleFunc(PATH_USUARIO_ID, servidor.DeletarUsuario).Methods(http.MethodDelete)

	fmt.Println("Escutando na porta 5000")
	log.Fatal(http.ListenAndServe(":5000", router))
}
