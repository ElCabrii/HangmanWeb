package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Printf("Server started at http://localhost:8080\n")
	player := Player{Username: "Joueur", Score: 0}
	http.HandleFunc("/", player.index)
	game := Game{Difficulty: 0, WordToGuess: "temp", Game: []string{"temp"}, WrongLetters: "", Mistakes: 0, GameOver: 0}
	http.HandleFunc("/play", game.play)
	handleDir()
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
