package HangmanController

import (
	"math/rand"
	"os"
	"strings"
)

func convertFileToSlice(filepath string) []string {
	var formattedList []string
	j := 0
	read, err := os.ReadFile(filepath)
	if err != nil {
		panic(err)
	}
	file := string(read)
	for i := 0; i < len(file); i++ {
		if rune(file[i]) == '\n' {
			formattedList = append(formattedList, file[j:i])
			j = i + 1
		}
	}
	return formattedList
}

func PickRandWord(levelChosen int) string {
	var wordToGuess string
	var file string
	if levelChosen == 1 {
		file = "HangmanWebpage/assets/wordsFiles/easy.txt"
	} else if levelChosen == 2 {
		file = "HangmanWebpage/assets/wordsFiles/medium.txt"
	} else if levelChosen == 3 {
		file = "HangmanWebpage/assets/wordsFiles/hard.txt"
	}
	slice := convertFileToSlice(file)
	lengthList := len(slice)
	wordToGuess = slice[rand.Intn(lengthList)]
	return strings.ToUpper(wordToGuess)
}
