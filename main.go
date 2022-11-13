package main

import (
	"bytes"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/urfave/cli"
)

var words = 4
var maxWordLength = 0
var dictionaryPath = "/usr/share/dict/words"

func main() {
	app := cli.NewApp()
	app.Name = "xpwd"
	app.Usage = "Suggest passwords in the style of XKCD"
	app.Action = func(c *cli.Context) {

		fmt.Println(strings.Join(RandomWords(words), " "))
	}

	app.Run(os.Args)
}

// RandomWords returns a space delimited list of words randomly selected from a dictionary
func RandomWords(wordCount int) []string {
	var randomWords []string
	for index := 0; index < wordCount; index++ {
		randomWord := RandomWord()
		randomWords = append(randomWords, randomWord)
	}

	return randomWords
}

// RandomWord selects a random word from newline-delimited list of dictionary words
func RandomWord() string {
	if maxWordLength > 0 && maxWordLength < 7 {
		log.Println("Word length must be at least 7.")
	}

	word := ""
	wordFound := false
	for wordFound == false {
		// seed rand function for each new word selected
		source := rand.NewSource(time.Now().UnixNano())
		seededRand := rand.New(source)

		totalWords := lineCount(dictionaryPath)
		// random line/word position in file
		randomLinePosition := strconv.Itoa(seededRand.Intn(totalWords))

		// for now, executing cli commands... until I learn more about file IO in golang

		// sed command "$ sed -n <linenumber>p filepath" will return contents of line
		word = strings.TrimSpace(commandOutput("sed", "-n", randomLinePosition+"p", dictionaryPath))
		wordFound = maxWordLength == 0 || len(word) <= maxWordLength
	}

	return word
}

// lineCount returns the number of lines in a given file
func lineCount(filePath string) int {
	// again, using cli commands for now, until I learn more about file IO in golang
	// wc command, with -l flag, returns the number of lines in a file
	wcOutput := strings.TrimSpace(commandOutput("wc", "-l", filePath))
	lineCount, _ := strconv.Atoi(strings.Split(wcOutput, " ")[0])

	return lineCount
}

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
