package main

import (
  "os"
  "github.com/codegangsta/cli"
)

func main() {
  app := cli.NewApp()
  app.Name = "cli"
  app.Usage = "a travis-ci command line client"
  app.Action = func(c *cli.Context) {
    println("hello world!")
  }

  app.Run(os.Args)
}

