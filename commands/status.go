package commands

import "github.com/codegangsta/cli"

var Status = cli.Command{
	Name:   "status",
	Usage:  "Show current status",
	Action: statusCmd,
}

func statusCmd(c *cli.Context) {

}
