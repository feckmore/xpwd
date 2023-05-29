package domain

import (
	"errors"
	"strings"
)

var ErrInvalidWordLength = errors.New("invalid word length")
var ErrNotEnoughWords = errors.New("insufficient available words")

// Generator is the interface that wraps the GenerateRandomWord method
type Generator interface {
	GenerateRandomWord() string
}

// Service is the interface that provides the Passphrase method
type Service interface {
	Passphrase(wordCount int) string
}

// service is a concrete implementation of the Service interface
type service struct {
	Generator
}

// New returns a new service implementation that fulfills the Service interface
func New(generator Generator) Service {
	return &service{generator}
}

// Passphrase returns a slice of random words
func (u *service) Passphrase(wordCount int) string {
	var randomWords []string
	for index := 0; index < wordCount; index++ {
		randomWord := u.Generator.GenerateRandomWord()
		randomWords = append(randomWords, randomWord)
	}

	return strings.Join(randomWords, " ")
}
