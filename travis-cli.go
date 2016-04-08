package main

import (
	"os"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.EnableBashCompletion = true
	app.Name = "cli"
	app.Usage = "a travis-ci command line client"
	app.Action = func(c *cli.Context) {
		println("Usage: travis-cli COMMAND ...\n \n Available commands: \n \n      login      logs you in!")

	}
	app.Commands = []cli.Command{
		{
			Name:    "login",
			Aliases: []string{"a"},
			Usage:   "add a task to the list",
			Action: func(c *cli.Context) {
				login()
			},
		},
	}

	app.Run(os.Args)
}
