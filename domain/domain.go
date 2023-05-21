package domain

import "errors"

var ErrInvalidWordLength = errors.New("invalid word length")

// Generator is the interface that wraps the GenerateRandomWord method
type Generator interface {
	GenerateRandomWord() string
}

// Service is the interface that provides the GenerateRandomWords method
type Service interface {
	GenerateRandomWords(wordCount int) []string
}

// RandomWordUsecase is a concrete implementation of the Service interface
type RandomWordUsecase struct {
	Generator
}

// New returns a new RandomWordUsecase that fulfills the Service interface
func New(generator Generator) Service {
	return &RandomWordUsecase{generator}
}

// GenerateRandomWords returns a slice of random words
func (u *RandomWordUsecase) GenerateRandomWords(wordCount int) []string {
	var randomWords []string
	for index := 0; index < wordCount; index++ {
		randomWord := u.Generator.GenerateRandomWord()
		randomWords = append(randomWords, randomWord)
	}

	return randomWords
}
