package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"

	"github.com/feckmore/xpwd/dictionary/oxford"
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
	}
	app.Action = func(c *cli.Context) {
		passphrase(c.Int("count"), c.Int("min"), c.Int("max"))
	}

	app.Run(os.Args)
}

func passphrase(wordCount, minWordLength, maxWordLength int) {
	// generator, err := mac.New(minWordLength, maxWordLength)
	generator, err := oxford.New(minWordLength, maxWordLength)

	if err != nil {
		log.Fatalf("Error creating dictionary: %v", err)
	}
	service := domain.New(generator)

	fmt.Println(service.Passphrase(wordCount))
}
