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
	Player       Player
	Difficulty   int
	WordToGuess  string
	Game         []string
	WrongLetters []string
	WrongGuesses int
}

func (player *Player) index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "HangmanWebpage/templates/index.html")

}

func (game *Game) play(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		err := r.ParseForm()
		if err != nil {
			fmt.Printf("Formulaire vide\n")
		}
		game.Player = Player{Username: r.FormValue("username"), Score: 0}
		game.Difficulty, _ = strconv.Atoi(r.FormValue("difficulty"))
		game.WordToGuess = HangmanController.PickRandWord(game.Difficulty)
		game.Game = HangmanController.InitGame(game.WordToGuess)
		template.Must(template.ParseFiles("HangmanWebpage/templates/play.html")).Execute(w, game)
	case "GET":
		err := r.ParseForm()
		if err != nil {
			fmt.Printf("Formulaire vide\n")
		}
		userInput := r.FormValue("userInput")
		game.Game = HangmanController.RefreshGame(userInput, game.WordToGuess, game.Game)
		fmt.Printf("Game: %v\n", game.Game)
		template.Must(template.ParseFiles("HangmanWebpage/templates/play.html")).Execute(w, game)
	}
}

func handleDir() {
	http.Handle("/HangmanWebpage/assets/", http.StripPrefix("/HangmanWebpage/assets", http.FileServer(http.Dir("HangmanWebpage/assets"))))
	http.Handle("/HangmanWebpage/templates/", http.StripPrefix("/HangmanWebpage/templates", http.FileServer(http.Dir("HangmanWebpage/templates"))))
	http.Handle("/HangmanWebpage/static/", http.StripPrefix("/HangmanWebpage/static", http.FileServer(http.Dir("HangmanWebpage/static"))))
}
