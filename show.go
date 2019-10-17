package main

import (
	"fmt"
	"strings"

	"github.com/urfave/cli"
)

func refineEnvVar(envVar string) string {
	return strings.Replace(envVar, "\n", "\\n", -1)
}

func CmdShow(c *cli.Context) (err error) {
	envMap, err := LoadS3()
	if err != nil {
		return
	}

	envVarName := c.Args().Get(0)
	if envVarName != "" {
		fmt.Print(refineEnvVar(envMap[envVarName]))
		return
	}

	for k, v := range envMap {
		if c.Bool("export") {
			fmt.Print("export ")
		}
		fmt.Printf("%s=\"%s\"\n", k, strings.Replace(refineEnvVar(v), "\"", "\\\"", -1))
	}
	return
}
