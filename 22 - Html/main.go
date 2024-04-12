package main

import (
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template

type usuario struct {
	Nome  string
	Email string
}

func main() {

	templates = template.Must(template.ParseGlob("*.html"))

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		u := usuario{"Louis", "joao@gmail.com"}
		templates.ExecuteTemplate(w, "home.html", u)
	})

	// sobe servidor
	log.Fatal(http.ListenAndServe(":5000", nil))
}
