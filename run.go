package main

import (
	"os"
	"os/exec"

	"github.com/urfave/cli"
)

func CmdRun(c *cli.Context) (err error) {
	if err = LoadS3(); err != nil {
		return
	}

	args := c.Args()
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
