package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/urfave/cli"
)

func refineEnvVar(envVar string) string {
	return strings.Replace(envVar, "\n", "\\n", -1)
}

func CmdShow(c *cli.Context) (err error) {
	if err = LoadS3(); err != nil {
		return
	}

	envVarName := c.Args().Get(0)
	if envVarName != "" {
		fmt.Print(refineEnvVar(os.Getenv(envVarName)))
		return
	}

	for _, env := range os.Environ() {
		s := strings.SplitN(env, "=", 2)
		if c.Bool("export") {
			fmt.Print("export ")
		}
		fmt.Printf("%s=\"%s\"\n", s[0], strings.Replace(refineEnvVar(s[1]), "\"", "\\\"", -1))
	}
	return
}
