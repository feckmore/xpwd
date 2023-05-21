package mac

import (
	"bytes"
	"log"
	"math/rand"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

// MAC_OS_LOCAL_SYSTEM_DICTIONARY_PATH is the path to the local system dictionary on Mac OS
// the dictionary is a list of words, one per line, that contains approx 235,000 words
const MAC_OS_LOCAL_SYSTEM_DICTIONARY_PATH = "/usr/share/dict/words"

// infrastructure package
type Dictionary struct {
	Path string
}

func New() *Dictionary {
	return &Dictionary{Path: MAC_OS_LOCAL_SYSTEM_DICTIONARY_PATH}
}

func (d *Dictionary) GenerateRandomWord(maxWordLength int) string {
	if maxWordLength > 0 && maxWordLength < 7 {
		log.Println("Word length must be at least 7.")
	}

	word := ""
	wordFound := false
	for !wordFound {
		// seed rand function for each new word selected
		source := rand.NewSource(time.Now().UnixNano())
		seededRand := rand.New(source)

		totalWords := lineCount(d.Path)
		// random line/word position in file
		randomLinePosition := strconv.Itoa(seededRand.Intn(totalWords))

		// for now, executing cli commands... until I learn more about file IO in golang

		// sed command "$ sed -n <linenumber>p filepath" will return contents of line
		word = strings.TrimSpace(commandOutput("sed", "-n", randomLinePosition+"p", d.Path))
		wordFound = maxWordLength == 0 || len(word) <= maxWordLength
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
	// again, using cli commands for now, until I learn more about file IO in golang
	// wc command, with -l flag, returns the number of lines in a file
	wcOutput := strings.TrimSpace(commandOutput("wc", "-l", filePath))
	lineCount, _ := strconv.Atoi(strings.Split(wcOutput, " ")[0])

	return lineCount
}
