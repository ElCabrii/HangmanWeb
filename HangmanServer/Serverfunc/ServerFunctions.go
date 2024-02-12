package Server

import (
	"Hangman-Web/HangmanController"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

type Player struct {
	Username  string
	Score     int
	HighScore int
}

type Game struct {
	Player            Player
	Difficulty        int
	WordToGuess       string
	Game              []string
	GameDisplay       string
	WrongLetters      string
	Mistakes          int
	RemainingMistakes int
	GameOver          int
	GameImage         string
	AlreadyFound      string
}

func (game *Game) Index(w http.ResponseWriter, r *http.Request) {
	execTmpl(w, "HangmanWebpage/templates/index.html", game)
}

func (game *Game) Play(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case "GET":
		game.resetGame()
		err := r.ParseForm()
		if err != nil {
			fmt.Printf("Formulaire vide\n")
		}

		game.Player.Username = r.FormValue("username")
		game.Difficulty, _ = strconv.Atoi(r.FormValue("difficulty"))
		game.WordToGuess = HangmanController.PickRandWord(game.Difficulty)
		game.Game = HangmanController.InitGame(game.WordToGuess)
		game.GameDisplay = HangmanController.PrintGame(game.Game)
		game.RemainingMistakes = 10
		game.GameImage = "HangmanWebpage/assets/HangmanBringToDeath/hangman0.png"
		execTmpl(w, "HangmanWebpage/templates/play.html", game)

	case "POST":

		if game.GameOver == 0 {

			err := r.ParseForm()
			if err != nil {
				fmt.Printf("Formulaire vide\n")
			}

			userInput := r.FormValue("userInput")
			game.Game, game.WrongLetters, game.AlreadyFound = HangmanController.RefreshGame(userInput, game.WordToGuess, game.Game, game.WrongLetters)
			game.Mistakes = len(game.WrongLetters) / 2
			game.RemainingMistakes = 10 - game.Mistakes
			game.GameImage = "HangmanWebpage/assets/HangmanBringToDeath/hangman" + strconv.Itoa(game.Mistakes) + ".png"
			game.GameDisplay = HangmanController.PrintGame(game.Game)
			game.GameOver = HangmanController.IsTheGameOver(game.GameDisplay, game.Mistakes, game.WordToGuess)

		}

		if game.GameOver == 1 {
			game.Player.Score = game.RemainingMistakes
			if game.Player.Score > game.Player.HighScore {
				game.Player.HighScore = game.Player.Score
			}
			execTmpl(w, "HangmanWebpage/templates/win.html", game)

		} else if game.GameOver == 2 {
			game.Player.Score = game.RemainingMistakes
			execTmpl(w, "HangmanWebpage/templates/lose.html", game)

		} else {

			execTmpl(w, "HangmanWebpage/templates/play.html", game)

		}
	}
}

func HandleDir() {
	http.Handle("/HangmanWebpage/assets/", http.StripPrefix("/HangmanWebpage/assets", http.FileServer(http.Dir("HangmanWebpage/assets"))))
	http.Handle("/HangmanWebpage/templates/", http.StripPrefix("/HangmanWebpage/templates", http.FileServer(http.Dir("HangmanWebpage/templates"))))
	http.Handle("/HangmanWebpage/static/", http.StripPrefix("/HangmanWebpage/static", http.FileServer(http.Dir("HangmanWebpage/static"))))
	http.Handle("HangmanWebpage/assets/HangmanBringToDeath", http.StripPrefix("HangmanWebpage/assets/HangmanBringToDeath", http.FileServer(http.Dir("HangmanWebpage/assets/HangmanBringToDeath"))))
}

func (game *Game) resetGame() {
	game.WordToGuess = ""
	game.Game = []string{}
	game.GameDisplay = ""
	game.WrongLetters = ""
	game.Mistakes = 0
	game.GameOver = 0
	game.Player.Score = 0
}

func execTmpl(w http.ResponseWriter, tmpl string, data interface{}) {
	err := template.Must(template.ParseFiles(tmpl)).Execute(w, data)
	if err != nil {
		fmt.Printf("Erreur d'execution du template\n")
	}
}
