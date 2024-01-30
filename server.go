package main

import (
	"fmt"
	"html/template"
	"net/http"
	hangman "server/src"
	"strconv"
)

type Page struct {
	Title string
	Body  []byte
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

func initHangman(difficulty int) string{
	wordToGuess := hangman.WordToGuess(difficulty)
	game := hangman.InitGame(wordToGuess)
	return game
}

func handleFunctions(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		tpl.ExecuteTemplate(w, "index.html", nil)
	case "/play":
		tpl.ExecuteTemplate(w, "play.html", nil)
	default:
		http.Error(w, "404 Not Found", 404)
	}
}



func main() {
	fmt.Printf("Starting server at port 8080...\n")
	http.HandleFunc("/", handleFunctions)
	http.HandleFunc("/play", handleFunctions)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	fmt.Printf("Server launched successfully !.\n")
	http.ListenAndServe(":8080", nil)
}
