package hangman

func WordToGuess(difficulty int)string {
	filePath := selectDifficulty(difficulty)
	return pickAWord(filePath)
}
