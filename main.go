package main

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/joho/godotenv"
	"github.com/urfave/cli"
)

func LoadS3() (err error) {
	sess := session.Must(session.NewSession())
	svc := s3.New(sess)

	ctx := context.Background()

	result, err := svc.GetObjectWithContext(ctx, &s3.GetObjectInput{
		Bucket: aws.String(os.Getenv("S3ENV_BUCKET_NAME")),
		Key:    aws.String(os.Getenv("S3ENV_KEY_NAME")),
	})
	if err != nil {
		return
	}

	f, err := ioutil.TempFile("", "dotenv")

	var r io.Reader
	r = result.Body
	if os.Getenv("S3ENV_BASE64ENCODE") == "y" {
		r = base64.NewDecoder(base64.StdEncoding, result.Body)
	}
	io.Copy(f, r)

	if err = godotenv.Load(f.Name()); err != nil {
		return
	}

	return
}

func runCmdStartPosition(args []string) (int, error) {
	for i, arg := range args {
		if arg == "run" {
			return i, nil
		}
	}
	return 0, errors.New("run command not found")
}

func handleRunCmd(args []string) {
	if len(args) <= 1 {
		fmt.Fprintln(os.Stderr, "Error:", "run command require command argument to run")
		os.Exit(1)
	}

	if err := CmdRun(args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
		return
	}
	os.Exit(0)
}

func main() {
	if idx, err := runCmdStartPosition(os.Args); err == nil {
		handleRunCmd(os.Args[idx:])
	}

	app := cli.NewApp()
	app.Name = "s3env"
	app.Usage = "Load environment variable from AWS S3"
	app.Version = Version
	app.Author = "sachaos"
	app.Email = "sakataku7@gmail.com"

	app.Commands = Commands

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}
