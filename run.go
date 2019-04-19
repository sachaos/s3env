package main

import (
	"os"
	"os/exec"
	"syscall"
)

func CmdRun(args []string) (err error) {
	if err = LoadS3(); err != nil {
		return
	}

	path, err := exec.LookPath(args[0])
	if err != nil {
		return err
	}

	return syscall.Exec(path, args, os.Environ())
}
