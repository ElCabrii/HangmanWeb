package main

import (
	"Hangman-Web/HangmanServer/Serverfunc"
	"fmt"
	"net/http"
)

func main() {
	fmt.Printf("Server starting at http://localhost:8080\n")
	http.HandleFunc("/", Server.Index)
	game := Server.Game{Difficulty: 0, WordToGuess: "temp", Game: []string{"temp"}, WrongLetters: "", Mistakes: 0, GameOver: 0}
	http.HandleFunc("/play", game.Play)
	Server.HandleDir()
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Server failed to start\n")
	} else {
		fmt.Printf("Server started successfully\n")
	}
}
