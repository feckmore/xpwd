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
	}
	app.Action = func(c *cli.Context) {
		count := c.Int("count")
		passphrase(count, 6, 11)
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
