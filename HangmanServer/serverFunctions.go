package main

import (
	"Hangman-Web/HangmanController"
	"fmt"
	"net/http"
	"strconv"
)

type Player struct {
	Username string
	Score    int
}

type Game struct {
	Difficulty  int
	WordToGuess string
	Game        string
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

func (game *Game) play(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "HangmanWebpage/templates/play.html")
	case "POST":
		err := r.ParseForm()
		if err != nil {
			return
		}
		game.Difficulty, _ = strconv.Atoi(r.FormValue("difficulty"))
		game.WordToGuess = HangmanController.PickRandWord(game.Difficulty)
		fmt.Printf(game.WordToGuess)
	}
}

func handleDir() {
	http.Handle("/HangmanWebpage/assets/", http.StripPrefix("/HangmanWebpage/assets", http.FileServer(http.Dir("HangmanWebpage/assets"))))
	http.Handle("/HangmanWebpage/templates/", http.StripPrefix("/HangmanWebpage/templates", http.FileServer(http.Dir("HangmanWebpage/templates"))))
	http.Handle("/HangmanWebpage/static/", http.StripPrefix("/HangmanWebpage/static", http.FileServer(http.Dir("HangmanWebpage/static"))))
}
