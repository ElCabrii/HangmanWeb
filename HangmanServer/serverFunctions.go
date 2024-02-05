package main

import (
	"Hangman-Web/HangmanController"
	"fmt"
	"net/http"
	"strconv"
	"text/template"
)

type Player struct {
	Username string
	Score    int
}

type Game struct {
	Player      Player
	Difficulty  int
	WordToGuess string
	Game        string
}

func (player *Player) index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "HangmanWebpage/templates/index.html")

}

func (game *Game) play(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		return
	}
	game.Player = Player{Username: r.FormValue("username"), Score: 0}
	difficulty := r.FormValue("difficulty")
	game.Difficulty, _ = strconv.Atoi(difficulty)
	game.WordToGuess = HangmanController.PickRandWord(game.Difficulty)
	fmt.Printf("Word to guess: %s\n", game.WordToGuess)
	fmt.Printf("Difficulty: %d\n", game.Difficulty)
	fmt.Printf("Player: %s\n", game.Player.Username)
	template.Must(template.ParseFiles("HangmanWebpage/templates/play.html")).Execute(w, game)
}

func handleDir() {
	http.Handle("/HangmanWebpage/assets/", http.StripPrefix("/HangmanWebpage/assets", http.FileServer(http.Dir("HangmanWebpage/assets"))))
	http.Handle("/HangmanWebpage/templates/", http.StripPrefix("/HangmanWebpage/templates", http.FileServer(http.Dir("HangmanWebpage/templates"))))
	http.Handle("/HangmanWebpage/static/", http.StripPrefix("/HangmanWebpage/static", http.FileServer(http.Dir("HangmanWebpage/static"))))
}
