package main

import (
	"net/http"
)

type Player struct {
	Username string
	Score    int
}

type Game struct {
	Difficulty int
}

func (player *Player) index(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "HangmanWebpage/templates/index.html")
	case "POST":
		err := r.ParseForm()
		if err != nil {
			return
		}
		player.Username = r.FormValue("username")
	}
}

func play(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "HangmanWebpage/templates/play.html")
}

func handleDir() {
	http.Handle("/HangmanWebpage/assets/", http.StripPrefix("/HangmanWebpage/assets", http.FileServer(http.Dir("HangmanWebpage/assets"))))
	http.Handle("/HangmanWebpage/templates/", http.StripPrefix("/HangmanWebpage/templates", http.FileServer(http.Dir("HangmanWebpage/templates"))))
	http.Handle("/HangmanWebpage/static/", http.StripPrefix("/HangmanWebpage/static", http.FileServer(http.Dir("HangmanWebpage/static"))))
}
