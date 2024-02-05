package main

import (
	"html/template"
	"net/http"
)

type Page struct {
	Title string
	Body  []byte
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/play", play)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		// Logique pour changer le contenu
		tmpl := template.Must(template.ParseFiles("chemin_vers_votre_fichier_html"))
		tmpl.Execute(w, "Nouveau contenu")
	} else {
		tmpl := template.Must(template.ParseFiles("chemin_vers_votre_fichier_html"))
		tmpl.Execute(w, "Contenu initial")
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.html", nil)
}

func play(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "play.html", nil)
}
