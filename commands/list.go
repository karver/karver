package commands

import "github.com/codegangsta/cli"

var List = cli.Command{
	Name:   "list",
	Usage:  "List migrations",
	Action: listCmd,
}

func listCmd(c *cli.Context) {

}
