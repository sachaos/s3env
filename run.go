package main

import (
	"os"
	"os/exec"
	"syscall"
)

func CmdRun(args []string) (err error) {
	envMap, err := LoadS3()
	if err != nil {
		return
	}

	for k, v := range envMap {
		os.Setenv(k, v)
	}

	path, err := exec.LookPath(args[0])
	if err != nil {
		return err
	}

	return syscall.Exec(path, args, os.Environ())
}
