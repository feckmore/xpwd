package oxford

import (
	"math/rand"
	"time"

	"github.com/feckmore/xpwd/domain"
)

// dictionary is a concrete implementation of the Generator interface
type dictionary struct {
	maxIndex int
	minIndex int
}

func New(minGeneratedWordLength, maxGeneratedWordLength int) (domain.Generator, error) {
	if maxGeneratedWordLength > 0 && minGeneratedWordLength > maxGeneratedWordLength {
		return nil, domain.ErrInvalidWordLength
	}

	if maxGeneratedWordLength < 0 {
		return nil, domain.ErrInvalidWordLength
	}

	maxPossibleWordLength := len(wordLengthIndices) - 1
	if minGeneratedWordLength > maxPossibleWordLength || maxGeneratedWordLength > maxPossibleWordLength {
		return nil, domain.ErrInvalidWordLength
	}

	minIndex := 0
	maxIndex := len(words) - 1

	minIndex = wordLengthIndices[minGeneratedWordLength]
	if maxGeneratedWordLength > 0 && maxGeneratedWordLength < maxPossibleWordLength {
		maxIndex = wordLengthIndices[maxGeneratedWordLength+1]
	}

	availableWords := maxIndex - minIndex + 1
	if availableWords < 500 {
		return nil, domain.ErrNotEnoughWords
	}

	return &dictionary{
		maxIndex: maxIndex,
		minIndex: minIndex,
	}, nil
}

// GenerateRandomWord returns a random word from the local system dictionary
func (d *dictionary) GenerateRandomWord() string {
	word := Word{}

	// seed rand function for each new word selected
	source := rand.NewSource(time.Now().UnixNano())
	seededRand := rand.New(source)

	// select ramdom word from dictionary within the min and max index
	availableWords := d.maxIndex - d.minIndex + 1
	index := d.minIndex + seededRand.Intn(availableWords)
	word = words[index]

	return word.Word
}
