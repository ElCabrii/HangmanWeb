package HangmanController

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

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
	wordToGuess = slice[rand.Intn(sliceLength)]
	return strings.ToUpper(wordToGuess)
}

func InitGame(wordToGuess string) []string {
	var game []string
	for i := 0; i < len(wordToGuess); i++ {
		game = append(game, "_")
	}
	return game
}

func RefreshGame(userInput string, wordToGuess string, game []string, wrongLetters string) ([]string, string, string) {
	userInput = strings.ToUpper(userInput)
	found := 0
	alreadyFound := ""

	for i := 0; i < len(wordToGuess); i++ {
		if game[i] == userInput {
			found = 1
		} else if string(wordToGuess[i]) == userInput {
			game[i] = userInput
			found = 2
		}
	}

	if found == 0 {
		for i := 0; i < len(wrongLetters); i++ {
			if string(wrongLetters[i]) == userInput {
				found = 1
			}
		}
		if found == 0 {
			wrongLetters = wrongLetters + " " + userInput
		}
	}
	if found == 1 || found == 3 {
		alreadyFound = "La lettre " + userInput + " a déjà été testée"
	}
	return game, wrongLetters, alreadyFound
}

func PrintGame(game []string) string {
	return strings.Join(game, "")
}

func IsTheGameOver(game string, mistakes int, wordToGuess string) int {
	if game == wordToGuess {
		return 1
	} else if mistakes == 10 {
		return 2
	} else {
		return 0
	}
}
