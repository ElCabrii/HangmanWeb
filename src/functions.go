package hangman

import (
	"os"
	"math/rand"
	"strings"
)

func selectDifficulty(difficulty int) string {
	switch difficulty {
	case 0:
		return "assets\\wordsFiles\\easy.txt"
	case 2:
		return "assets\\wordsFiles\\medium.txt"
	case 3:
		return "assets\\wordsFiles\\hard.txt"
	default:
		return "assets\\wordsFiles\\easy.txt"
	}	
}

func pickAWord(filePath string) string {
	var wordToGuess string
	var formattedList []string
	j := 0
	read, err := os.ReadFile(filePath)
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
	lengthList := len(formattedList)
	wordToGuess = formattedList[rand.Intn(lengthList)]
	return strings.ToUpper(wordToGuess)
}

func InitGame(wordToGuess string)string{
	var game string
	for i := 0; i < len(wordToGuess)-1; i++ {
		game += "_ "
	}
	return game
}