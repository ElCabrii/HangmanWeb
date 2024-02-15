package Server

import (
	"Hangman-Web/HangmanController"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

// Player is used to store the player's username and scores
type Player struct {
	Username  string
	Score     int
	HighScore int
}

// Game is used to store the game's data
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

// Index is used to display the index page
func (game *Game) Index(w http.ResponseWriter, r *http.Request) {
	execTmpl(w, "HangmanWebpage/templates/index.html", game)
}

// Play is used to display the game page
func (game *Game) Play(w http.ResponseWriter, r *http.Request) {

	//Check if the request is a GET or a POST
	switch r.Method {

	//If the request is a GET, it means the request comes from the /index form, so the game is reset and initialized and the play page is displayed
	case "GET":

		game.resetGame()
		err := r.ParseForm()
		if err != nil {
			fmt.Printf("Formulaire vide\n")
		}

		//The player's username and the difficulty chosen are stored in the game struct
		game.Player.Username = r.FormValue("username")
		game.Difficulty, _ = strconv.Atoi(r.FormValue("difficulty"))

		//The word to guess is picked and the game is initialized
		game.WordToGuess = HangmanController.PickRandWord(game.Difficulty)
		game.Game = HangmanController.InitGame(game.WordToGuess)

		//Converting the game slice into a string to display it
		game.GameDisplay = HangmanController.PrintGame(game.Game)
		game.RemainingMistakes = 10

		//The game image is set to the first image. It will be updated depending on the mistakes made by the player and used in the play.html template
		game.GameImage = "HangmanWebpage/assets/HangmanBringToDeath/hangman0.png"

		//The play page is displayed with the game data
		execTmpl(w, "HangmanWebpage/templates/play.html", game)

	//If the request is a POST, it means the request comes from the /play form. It means the game is ongoing.
	case "POST":

		//If the game is not over, the game is refreshed and the play page is displayed
		if game.GameOver == 0 {

			err := r.ParseForm()
			if err != nil {
				fmt.Printf("Formulaire vide\n")
			}

			//The letter submitted is stored in userInput
			userInput := r.FormValue("userInput")

			//The letter is tested and the game is refreshed
			game.Game, game.WrongLetters, game.AlreadyFound = HangmanController.RefreshGame(userInput, game.WordToGuess, game.Game, game.WrongLetters)

			//The mistakes are recalculated
			game.Mistakes = len(game.WrongLetters)
			game.RemainingMistakes = 10 - game.Mistakes

			//The game image is updated
			game.GameImage = "HangmanWebpage/assets/HangmanBringToDeath/hangman" + strconv.Itoa(game.Mistakes) + ".png"

			//The game to display is updated
			game.GameDisplay = HangmanController.PrintGame(game.Game)

			//The game is checked to see if it's over
			game.GameOver = HangmanController.IsTheGameOver(game.GameDisplay, game.Mistakes, game.WordToGuess)
		}

		//If IsTheGameOver returns 1, the player wins and the win page is displayed.
		if game.GameOver == 1 {
			//The player's score is updated and the highscore is updated if the player's score is higher than the highscore
			game.Player.Score = game.RemainingMistakes
			if game.Player.Score > game.Player.HighScore {
				game.Player.HighScore = game.Player.Score
			}
			execTmpl(w, "HangmanWebpage/templates/win.html", game)

			//If IsTheGameOver returns 2, the player loses and the loss page is displayed.
		} else if game.GameOver == 2 {
			execTmpl(w, "HangmanWebpage/templates/lose.html", game)

			//Else, it means the game is not over, the play page is displayed again
		} else {
			execTmpl(w, "HangmanWebpage/templates/play.html", game)
		}
	}
}

// HandleDir is used to handle the directories used in the webpages
func HandleDir() {
	http.Handle("/HangmanWebpage/assets/", http.StripPrefix("/HangmanWebpage/assets", http.FileServer(http.Dir("HangmanWebpage/assets"))))
	http.Handle("/HangmanWebpage/templates/", http.StripPrefix("/HangmanWebpage/templates", http.FileServer(http.Dir("HangmanWebpage/templates"))))
	http.Handle("/HangmanWebpage/static/", http.StripPrefix("/HangmanWebpage/static", http.FileServer(http.Dir("HangmanWebpage/static"))))
	http.Handle("HangmanWebpage/assets/HangmanBringToDeath", http.StripPrefix("HangmanWebpage/assets/HangmanBringToDeath", http.FileServer(http.Dir("HangmanWebpage/assets/HangmanBringToDeath"))))
}

// resetGame is used to reset the game to an initial state
func (game *Game) resetGame() {
	game.WordToGuess = ""
	game.Game = []string{}
	game.GameDisplay = ""
	game.WrongLetters = ""
	game.Mistakes = 0
	game.GameOver = 0
	game.Player.Score = 0
}

// execTmpl is used to execute the templates and handle the errors
func execTmpl(w http.ResponseWriter, tmpl string, data interface{}) {
	err := template.Must(template.ParseFiles(tmpl)).Execute(w, data)
	if err != nil {
		fmt.Printf("Erreur d'execution du template\n")
	}
}
