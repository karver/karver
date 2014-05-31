package commands

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/karver/karver/migrations"
	"os"
)

var Status = cli.Command{
	Name:   "status",
	Usage:  "Show current status",
	Action: statusCmd,
}

func statusCmd(c *cli.Context) {
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

	if last := migrations.Last(list, timestamp); last != nil {
		fmt.Printf("Last migration executed: %s\n", last.Name)
	} else {
		fmt.Printf("No migration executed\n")
	}

	fmt.Printf("Pending migrations: %d\n", len(migrations.Pending(list, timestamp)))
}
