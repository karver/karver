package commands

import (
	"fmt"
	"log"

	"github.com/codegangsta/cli"
	"github.com/karver/karver/migrations"
)

var Run = cli.Command{
	Name:   "run",
	Usage:  "Run the migrations",
	Action: runCmd,
}

func runCmd(c *cli.Context) {
	target := "/"
	m := &migrations.Migration{
		Timestamp: "foo",
		Name:      "run",
		Path:      "/Users/salvador/devel/golang/src/github.com/karver/karver/run",
	}

	stdout, stderr, err := m.Run(target)

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(stdout)
	fmt.Println(stderr)

}
