package main

import (
	"Hangman-Web/HangmanServer/Serverfunc"
	"fmt"
	"net/http"
)

func main() {
	fmt.Printf("Server starting at http://localhost:8080\n")
	player := Server.Player{Username: "Invité", Score: 0, HighScore: 0}
	game := Server.Game{Difficulty: 2, WordToGuess: "", Game: []string{""}, WrongLetters: "", Mistakes: 0, GameOver: 0, Player: player}
	http.HandleFunc("/", game.Index)
	http.HandleFunc("/play", game.Play)
	Server.HandleDir()
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Server failed to start\n")
	}
}
