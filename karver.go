package main

import (
	"os"

	"github.com/codegangsta/cli"
	"github.com/karver/karver/commands"
)

const APP_VER = "0.0.1"

func main() {
	app := cli.NewApp()
	app.Name = "Karver"
	app.Usage = "Run filesystem migrations"
	app.Author = "Karver"
	app.Version = APP_VER
	app.Commands = []cli.Command{
		commands.List,
		commands.Run,
		commands.Status,
	}
	app.Flags = []cli.Flag{
		cli.StringFlag{"target", "/", "Target of the migrations. Will work as the root path"},
	}
	app.Run(os.Args)
}
