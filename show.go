package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/urfave/cli"
)

func CmdShow(c *cli.Context) (err error) {
	if err = LoadS3(); err != nil {
		return
	}

	envVarName := c.Args().Get(0)
	if envVarName != "" {
		fmt.Print(os.Getenv(envVarName))
		return
	}

	for _, env := range os.Environ() {
		s := strings.Split(env, "=")
		if c.Bool("export") {
			fmt.Print("export ")
		}
		fmt.Printf("%s=\"%s\"\n", s[0], s[1])
	}
	return
}
