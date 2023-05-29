package oxford

import (
	"math/rand"
	"time"

	"github.com/feckmore/xpwd/domain"
)

// dictionary is a concrete implementation of the Generator interface
type dictionary struct {
	MaxWordLength int
	MinWordLength int
}

func New(minWordLength, maxWordLength int) (domain.Generator, error) {
	if minWordLength < 3 {
		minWordLength = 3
	}

	if maxWordLength > 0 && minWordLength > maxWordLength {
		return nil, domain.ErrInvalidWordLength
	}

	return &dictionary{
		MaxWordLength: maxWordLength,
		MinWordLength: minWordLength,
	}, nil
}

// GenerateRandomWord returns a random word from the local system dictionary
func (d *dictionary) GenerateRandomWord() string {
	words := append(append(append(append(append([]Word{}, wordsA1...), wordsA2...), wordsB1...), wordsB2...), wordsC1...)
	word := Word{}
	wordFound := false
	for !wordFound {
		// seed rand function for each new word selected
		source := rand.NewSource(time.Now().UnixNano())
		seededRand := rand.New(source)
		// choose a random line number to select a word from the dictionary
		// randomLinePosition := strconv.Itoa(seededRand.Intn(totalWords))
		index := seededRand.Intn(len(words))
		word = words[index]

		wordFound = len(word.Word) >= d.MinWordLength && (d.MaxWordLength == 0 || len(word.Word) <= d.MaxWordLength)
	}

	return word.Word
}

// internal
