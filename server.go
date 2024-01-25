package main

import (
	"html/template"
	"net/http"
)

type Page struct {
	Valeur string
}

var templates = template.Must(template.ParseFiles("index.html"))

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	p := Page{Valeur: "webpage"}
	err := templates.ExecuteTemplate(w, "index.html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}