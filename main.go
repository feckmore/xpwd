package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/urfave/cli"

	"github.com/feckmore/xpwd/dictionary/system/mac"
	"github.com/feckmore/xpwd/domain"
)

func main() {
	app := cli.NewApp()
	app.Name = "xpwd"
	app.Usage = "Suggest passwords in the style of XKCD"
	app.Action = func(c *cli.Context) {
		dictionary := mac.New()
		usecase := &domain.RandomWordUsecase{Generator: dictionary}

		fmt.Println(strings.Join(usecase.GenerateRandomWords(4, 0), " "))
	}

	app.Run(os.Args)
}
