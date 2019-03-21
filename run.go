package main

import (
	"os"
	"os/exec"
	"os/signal"
)

func CmdRun(args []string) (err error) {
	if err = LoadS3(); err != nil {
		return
	}

	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	signals := make(chan os.Signal, 1)
	signal.Notify(signals)

	err = cmd.Start()
	if err != nil {
		return err
	}

	go func() {
		s := <- signals
		err := cmd.Process.Signal(s)
		if err != nil {
			panic(err)
		}
	}()

	return cmd.Wait()
}
