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

	bc := core.NewBlockchain()
	UTXOSet := core.UTXOSet{bc}
	defer bc.DB.Close()

	tx := core.NewUTXOTransaction(from, to, amount, &UTXOSet)
	cbTx := core.NewCoinbaseTX(from, "")
	txs := []*core.Transaction{cbTx, tx}

	newBlock := bc.MineBlock(txs)
	UTXOSet.Update(newBlock)
	fmt.Println("Success!")
}
