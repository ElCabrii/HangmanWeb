package HangmanController

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

// convertFileToSlice is used to convert a wordlist file into a slice of words, taking the file path as an argument
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

// PickRandWord is used to pick a random word from a wordlist file, taking the level chosen as an argument
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

// InitGame is used to create the unfounded word, using the length of the word to guess
func InitGame(wordToGuess string) []string {
	var game []string
	for i := 0; i < len(wordToGuess); i++ {
		game = append(game, "_")
	}
	return game
}

// RefreshGame is used to refresh the game with the user input. It acts depending on wether the user input is in the word to guess or not
func RefreshGame(userInput string, wordToGuess string, game []string, wrongLetters string) ([]string, string, string) {
	userInput = strings.ToUpper(userInput)
	found := 0
	alreadyFound := ""
	//If the user input is in the word to guess, the game is updated
	for i := 0; i < len(wordToGuess); i++ {
		if game[i] == userInput {
			found = 1
		} else if string(wordToGuess[i]) == userInput {
			game[i] = userInput
			found = 2
		}
	}
	//If the user input is not in the word to guess
	if found == 0 {
		//We check if the user input is in the wrong letters list
		for i := 0; i < len(wrongLetters); i++ {
			if string(wrongLetters[i]) == userInput {
				found = 1
			}
		}
		//If not, we add it to the wrong letters list
		if found == 0 {
			wrongLetters = wrongLetters + " " + userInput
		}
	}
	//If the user input is already in the game or in the wrong letters list, we return a message to inform the user
	if found == 1 || found == 3 {
		alreadyFound = "La lettre " + userInput + " a déjà été testée"
	}
	return game, wrongLetters, alreadyFound
}

// PrintGame is used to convert the game slice into a string
func PrintGame(game []string) string {
	return strings.Join(game, "")
}

// IsTheGameOver is used to check if the game is over depending on the amount of mistakes and the game state
func IsTheGameOver(game string, mistakes int, wordToGuess string) int {
	if game == wordToGuess {
		return 1
	} else if mistakes == 10 {
		return 2
	} else {
		return 0
	}
}
