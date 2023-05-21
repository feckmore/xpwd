package domain

type Generator interface {
	GenerateRandomWord(maxWordLength int) string
}

// domain package
type Service interface {
	GenerateRandomWords(wordCount, maxWordLength int) []string
}

// usecase package
type RandomWordUsecase struct {
	Generator
}

func New(generator Generator) Service {
	return &RandomWordUsecase{generator}
}

func (u *RandomWordUsecase) GenerateRandomWords(wordCount, maxWordLength int) []string {
	var randomWords []string
	for index := 0; index < wordCount; index++ {
		randomWord := u.Generator.GenerateRandomWord(maxWordLength)
		randomWords = append(randomWords, randomWord)
	}

	return randomWords
}
