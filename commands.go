package main

import (
	"github.com/urfave/cli"
)

var Commands = []cli.Command{
	{
		Name:   "show",
		Usage:  "Display environment variables",
		Action: CmdShow,
		Flags: []cli.Flag{
			cli.BoolFlag{
				Name: "export",
			},
		},
	},
	{
		Name:   "run",
		Usage:  "Execute command with loaded environment variables",
		Action: CmdRun,
	},
}
