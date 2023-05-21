package mac

import (
	"bytes"
	"log"
	"math/rand"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/feckmore/xpwd/domain"
)

// MAC_OS_LOCAL_SYSTEM_DICTIONARY_PATH is the path to the local system dictionary on Mac OS
// the dictionary is a list of words, one per line, that contains approx 235,000 words
const MAC_OS_LOCAL_SYSTEM_DICTIONARY_PATH = "/usr/share/dict/words"

// Dictionary is a concrete implementation of the Generator interface
type Dictionary struct {
	MaxWordLength int
	MinWordLength int
	Path          string
}

// New returns a new Dictionary that fulfills the Generator interface
func New(minWordLength, maxWordLength int) (domain.Generator, error) {
	if minWordLength < 3 {
		minWordLength = 3
	}

	if maxWordLength > 0 && minWordLength > maxWordLength {
		return nil, domain.ErrInvalidWordLength
	}

	return &Dictionary{
		MaxWordLength: maxWordLength,
		MinWordLength: minWordLength,
		Path:          MAC_OS_LOCAL_SYSTEM_DICTIONARY_PATH,
	}, nil
}

// GenerateRandomWord returns a random word from the local system dictionary
func (d *Dictionary) GenerateRandomWord() string {
	word := ""
	wordFound := false
	for !wordFound {
		// number of words in the dictionary
		totalWords := lineCount(d.Path)

		// seed rand function for each new word selected
		source := rand.NewSource(time.Now().UnixNano())
		seededRand := rand.New(source)
		// choose a random line number to select a word from the dictionary
		randomLinePosition := strconv.Itoa(seededRand.Intn(totalWords))

		// sed command "$ sed -n <linenumber>p filepath" will return contents of a line in a file
		word = strings.TrimSpace(commandOutput("sed", "-n", randomLinePosition+"p", d.Path))

		wordFound = len(word) >= d.MinWordLength && (d.MaxWordLength == 0 || len(word) <= d.MaxWordLength)
	}

	return word
}

// internal

// commandOutput executes the named command with arguments, returning Stdout
func commandOutput(name string, arg ...string) string {
	var out bytes.Buffer

	cmd := exec.Command(name, arg...)
	cmd.Stdout = &out

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}

	return out.String()
}

// lineCount returns the number of lines in a given file
func lineCount(filePath string) int {
	// wc command, with -l flag, returns the number of lines in a file
	wcOutput := strings.TrimSpace(commandOutput("wc", "-l", filePath))
	lineCount, _ := strconv.Atoi(strings.Split(wcOutput, " ")[0])

	return lineCount
}
