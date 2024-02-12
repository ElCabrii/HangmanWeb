package HangmanController

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

func randInt(max int) int {
	return rand.Intn(max)
}

func convertFileToSlice(filepath string) []string {
	var formattedList []string
	j := 0
	read, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Printf("Fichier introuvable")
	}
	file := string(read)
	for i := 0; i < len(file); i++ {
		if string(file[i]) == "\n" {
			formattedList = append(formattedList, file[j:i-1])
			j = i + 1
		}
	}
	return formattedList
}

func PickRandWord(levelChosen int) string {
	var wordToGuess string
	var file string
	if levelChosen == 0 {
		fmt.Printf("Niveau non choisi")
	} else if levelChosen == 1 {
		file = "HangmanWebpage/assets/Wordlists/easy.txt"
	} else if levelChosen == 2 {
		file = "HangmanWebpage/assets/Wordlists/medium.txt"
	} else if levelChosen == 3 {
		file = "HangmanWebpage/assets/Wordlists/hard.txt"
	}
	slice := convertFileToSlice(file)
	sliceLength := len(slice)
	wordToGuess = slice[randInt(sliceLength)]
	return strings.ToUpper(wordToGuess)
}

func InitGame(wordToGuess string) []string {
	var game []string
	for i := 0; i < len(wordToGuess); i++ {
		game = append(game, "_")
	}
	return game
}

func RefreshGame(userInput string, wordToGuess string, game []string, wrongLetters string) ([]string, string) {
	userInput = strings.ToUpper(userInput)
	var found bool
	for i := 0; i < len(wordToGuess); i++ {
		if string(wordToGuess[i]) == userInput {
			game[i] = userInput
			found = true
		}
	}
	if found == false {
		wrongLetters = wrongLetters + " " + userInput
	}
	return game, wrongLetters
}

func PrintGame(game []string) string {
	return strings.Join(game, "")
}

func IsTheGameOver(game string, mistakes int, wordToGuess string) int {
	if game == wordToGuess {
		fmt.Printf("Partie terminée\n")
		return 1
	} else if mistakes == 10 {
		return 2
	} else {
		return 0
	}
}
