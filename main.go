package main

import (
	"go_chain/cli"
	"os"
)

func main() {
	defer os.Exit(0)
	cli := cli.CLI{}
	cli.Run()
}
