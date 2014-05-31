package commands

import "github.com/codegangsta/cli"

var Run = cli.Command{
	Name:   "run",
	Usage:  "Run the migrations",
	Action: runCmd,
}

func runCmd(c *cli.Context) {

}
