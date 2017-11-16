package main

import (
	"github.com/sasaxie/go-blockchain/cli"
	"github.com/sasaxie/go-blockchain/core"
)

func main() {
	bc := core.NewBlockchain()
	defer bc.DB.Close()

	cli := cli.CLI{bc}
	cli.Run()
}
