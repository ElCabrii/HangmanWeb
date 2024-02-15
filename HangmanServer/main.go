package main

import (
	"Hangman-Web/HangmanServer/Serverfunc"
	"fmt"
	"net/http"
)

func main() {
	fmt.Printf("Server starting at http://localhost:8080\n")

	//Create a player and a game
	player := Server.Player{Username: "Invité", Score: 0, HighScore: 0}
	game := Server.Game{Difficulty: 2, Player: player}

	//Handle the routes
	http.HandleFunc("/", game.Index)
	http.HandleFunc("/play", game.Play)

	//Handle the assets
	Server.HandleDir()

	//Start the server
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Server failed to start\n")
	}
}
