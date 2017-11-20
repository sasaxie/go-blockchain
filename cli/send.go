package cli

import (
	"fmt"
	"github.com/sasaxie/go-blockchain/core"
	"log"
)

func (cli *CLI) send(from, to string, amount int) {
	if !core.ValidateAddress(from) {
		log.Panic("ERROR: Sender address is not valid")
	}
	if !core.ValidateAddress(to) {
		log.Panic("ERROR: Recipient address is not valid")
	}

	bc := core.NewBlockchain(from)
	defer bc.DB.Close()

	tx := core.NewUTXOTransaction(from, to, amount, bc)
	bc.MineBlock([]*core.Transaction{tx})
	fmt.Println("Success!")
}
