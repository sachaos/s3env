package main

import (
	"os"
	"os/exec"
)

func CmdRun(args []string) (err error) {
	if err = LoadS3(); err != nil {
		return
	}

	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
