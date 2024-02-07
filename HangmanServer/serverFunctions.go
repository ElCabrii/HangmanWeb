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
	GameDisplay  string
	WrongLetters string
	Mistakes     int
	GameOver     int
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
		game.GameDisplay = HangmanController.PrintGame(game.Game)
		template.Must(template.ParseFiles("HangmanWebpage/templates/play.html")).Execute(w, game)

	case "GET":

		if game.GameOver == 0 {

			err := r.ParseForm()
			if err != nil {
				fmt.Printf("Formulaire vide\n")
			}

			userInput := r.FormValue("userInput")

			game.Game, game.WrongLetters = HangmanController.RefreshGame(userInput, game.WordToGuess, game.Game, game.WrongLetters)
			game.Mistakes = len(game.WrongLetters)
			game.GameDisplay = HangmanController.PrintGame(game.Game)
			game.GameOver = HangmanController.IsTheGameOver(game.GameDisplay, game.Mistakes, game.WordToGuess)

		}

		if game.GameOver == 1 {

			template.Must(template.ParseFiles("HangmanWebpage/templates/win.html")).Execute(w, game)
			resetGame(game)

		} else if game.GameOver == 2 {

			template.Must(template.ParseFiles("HangmanWebpage/templates/lose.html")).Execute(w, game)
			resetGame(game)

		} else {

			template.Must(template.ParseFiles("HangmanWebpage/templates/play.html")).Execute(w, game)

		}
	}
}

func handleDir() {
	http.Handle("/HangmanWebpage/assets/", http.StripPrefix("/HangmanWebpage/assets", http.FileServer(http.Dir("HangmanWebpage/assets"))))
	http.Handle("/HangmanWebpage/templates/", http.StripPrefix("/HangmanWebpage/templates", http.FileServer(http.Dir("HangmanWebpage/templates"))))
	http.Handle("/HangmanWebpage/static/", http.StripPrefix("/HangmanWebpage/static", http.FileServer(http.Dir("HangmanWebpage/static"))))
}

func resetGame(game *Game) {
	game.Difficulty = 0
	game.WordToGuess = ""
	game.Game = []string{}
	game.GameDisplay = ""
	game.WrongLetters = ""
	game.Mistakes = 0
	game.GameOver = 0
}
