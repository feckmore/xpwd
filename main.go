package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"

	"github.com/feckmore/xpwd/dictionary/oxford"
	"github.com/feckmore/xpwd/dictionary/system/mac"
	"github.com/feckmore/xpwd/domain"
)

func main() {
	app := cli.NewApp()
	app.Name = "xpwd"
	app.Usage = "Suggest passwords in the style of XKCD"
	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name:  "count, c",
			Value: 4,
			Usage: "number of words in the passphrase",
		},
		cli.IntFlag{
			Name:  "min, m",
			Value: 6,
			Usage: "minimum word length",
		},
		cli.IntFlag{
			Name:  "max, x",
			Value: 11,
			Usage: "maximum word length",
		},
		cli.StringFlag{
			Name:  "dictionary, d",
			Value: "oxford",
			Usage: "dictionary to use",
		},
	}
	app.Action = func(c *cli.Context) {
		passphrase(
			c.String("dictionary"),
			c.Int("count"),
			c.Int("min"),
			c.Int("max"),
		)
	}

	app.Run(os.Args)
}

func passphrase(dictionary string, wordCount, minWordLength, maxWordLength int) {
	dictionaries := map[string]func(int, int) (domain.Generator, error){
		"oxford": oxford.New,
		"mac":    mac.New,
	}
	generator, err := dictionaries[dictionary](minWordLength, maxWordLength)

	if err != nil {
		log.Fatalf("Error creating dictionary: %v", err)
	}
	service := domain.New(generator)

	fmt.Println(service.Passphrase(wordCount))
}
