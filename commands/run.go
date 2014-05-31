package commands

import (
	"github.com/codegangsta/cli"
	"github.com/karver/karver/migrations"
	"log"
	"os"
)

var Run = cli.Command{
	Name:   "run",
	Usage:  "Run the migrations",
	Action: runCmd,
}

func runCmd(c *cli.Context) {
	migrationsPath := c.GlobalString("migrations")
	targetPath := c.GlobalString("target")

	migrationsPath, err := migrations.AbsMigrationsPath(migrationsPath)
	if err != nil {
		log.Fatal("Error determining the migrations path: " + err.Error())
		os.Exit(1)
	}

	list, err := migrations.List(migrationsPath)
	if err != nil {
		log.Fatal("Error listing the available migrations: " + err.Error())
		os.Exit(1)
	}

	timestamp, err := migrations.CurrentTimestamp(targetPath)
	if err != nil {
		log.Fatal("Error reading the current timestamp: " + err.Error())
		os.Exit(1)
	}

	pending := migrations.Pending(list, timestamp)

	migrations.Run(pending, targetPath)
}
