package commands

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/karver/karver/migrations"
	"os"
)

var List = cli.Command{
	Name:   "list",
	Usage:  "List migrations",
	Action: listCmd,
}

func listCmd(c *cli.Context) {
	migrationsPath := c.GlobalString("migrations")
	targetPath := c.GlobalString("target")

	migrationsPath, err := migrations.AbsMigrationsPath(migrationsPath)
	if err != nil {
		fmt.Printf("Error determining the migrations path: %s\n", err.Error())
		os.Exit(1)
	}

	list, err := migrations.List(migrationsPath)
	if err != nil {
		fmt.Printf("Error listing the available migrations: %s\n", err.Error())
		os.Exit(1)
	}

	timestamp, err := migrations.CurrentTimestamp(targetPath)
	if err != nil {
		fmt.Printf("Error reading the current timestamp: %s\n", err.Error())
		os.Exit(1)
	}

	if len(list) == 0 {
		fmt.Println("No migrations found")
		os.Exit(0)
	}

	for _, migration := range list {
		var symbol string

		needs := migrations.NeedsToRun(timestamp, migration)

		if needs {
			symbol = "✘"
		} else {
			symbol = "✔"
		}

		fmt.Printf("%s - %s\n", symbol, migration.Name)
	}
}
