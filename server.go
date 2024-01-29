package main

import (
	"net/http"
	"html/template"
)

type Page struct {
	Title string
	Body []byte
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}


func home(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.html", nil)
}

func play(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "play.html", nil)
}

func launch() {
	http.HandleFunc("/", home)
	http.HandleFunc("/play", play)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	http.ListenAndServe(":8080", nil)
}

func main() {
	launch()

}
