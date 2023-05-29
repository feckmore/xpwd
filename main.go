package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/urfave/cli"

	"github.com/feckmore/xpwd/dictionary/oxford"
	"github.com/feckmore/xpwd/domain"
)

func main() {
	// generator, err := mac.New(4, 8)
	generator, err := oxford.New(6, 11)

	if err != nil {
		log.Fatalf("Error creating dictionary: %v", err)
	}
	usecase := domain.New(generator)

	app := cli.NewApp()
	app.Name = "xpwd"
	app.Usage = "Suggest passwords in the style of XKCD"
	app.Action = func(c *cli.Context) {
		fmt.Println(strings.Join(usecase.GenerateRandomWords(4), " "))
	}

	app.Run(os.Args)
}
