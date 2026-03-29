package main

import (
	"context"
	"os"

	"github.com/mekedron/ocr/internal/cli"
)

var version = "dev"

func main() {
	exitCode := cli.Execute(context.Background(), os.Args[1:], version, os.Stdout, os.Stderr)
	os.Exit(exitCode)
}
