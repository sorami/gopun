package main

import (
	"os"
    "fmt"
	"bufio"

	"github.com/codegangsta/cli"
	"github.com/ikawaha/kagome/tokenizer"
)

func main() {

	app := cli.NewApp()
	app.Name = Name
	app.Version = Version
	app.Author = "Sorami Hisamoto"
	app.Email = ""
	app.Usage = "Create Golang pun in Japanese."

	app.Flags = GlobalFlags
	app.Commands = Commands
	app.CommandNotFound = CommandNotFound

	app.Action = func(c *cli.Context) {
		scanner := bufio.NewScanner(os.Stdin)
		t := tokenizer.New()

		for scanner.Scan() {
			line := scanner.Text()

			tokens := t.Tokenize(line)
			for _, token := range tokens {
				if token.Class == tokenizer.DUMMY {
					continue
				}
				yomi := token.Features()[7]
				if yomi == "ゴラン" {
					fmt.Printf("Golang")
				} else if yomi == "ゴ" {
					fmt.Printf("Go")
				} else {
					fmt.Printf("%s", token.Surface)
				}
			}
			fmt.Printf("\n")
		}
	}

	app.Run(os.Args)
}
