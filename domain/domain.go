package domain

// domain package
type RandomWordGenerator interface {
	GenerateRandomWord(maxWordLength int) string
}

// usecase package
type RandomWordUsecase struct {
	Generator RandomWordGenerator
}

func (u *RandomWordUsecase) GenerateRandomWords(wordCount, maxWordLength int) []string {
	var randomWords []string
	for index := 0; index < wordCount; index++ {
		randomWord := u.Generator.GenerateRandomWord(maxWordLength)
		randomWords = append(randomWords, randomWord)
	}

	return randomWords
}
