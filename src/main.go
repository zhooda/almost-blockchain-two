package main

import (
	"os"

	"github.com/zhooda/almost-blockchain-two/cli"
)

func main() {
	defer os.Exit(0)
	cli := cli.CommandLine{}
	cli.Run()
}
