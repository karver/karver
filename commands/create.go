package commands

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/karver/karver/migrations"
	"os"
)

var Create = cli.Command{
	Name:   "create",
	Usage:  "Create migrations",
	Action: createCmd,
}

func createCmd(c *cli.Context) {
	migrationsPath := c.GlobalString("migrations")

	var title string

	if len(c.Args()) > 0 {
		title = c.Args()[0]
	} else {
		fmt.Println("No migration title provided on create!")
		os.Exit(1)
	}

	migrationsPath, err := migrations.AbsMigrationsPath(migrationsPath)
	if err != nil {
		fmt.Printf("Error determining the migrations path: %s\n", err.Error())
		os.Exit(1)
	}

	m, err := migrations.Create(title, migrationsPath)
	if err != nil {
		fmt.Printf("Error creating the migration file: %s\n", err.Error())
		os.Exit(1)
	}

	fmt.Printf("New migration: %s - %s\n", m.Name, m.Path)
}
