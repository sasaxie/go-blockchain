package cli

import (
	"fmt"
	"github.com/sasaxie/go-blockchain/core"
	"github.com/sasaxie/go-blockchain/utils"
	"log"
)

func (cli *CLI) getBalance(address string) {
	if !core.ValidateAddress(address) {
		log.Panic("ERROR: Address is not valid")
	}
	bc := core.NewBlockchain()
	UTXOSet := core.UTXOSet{bc}
	defer bc.DB.Close()

	balance := 0
	pubKeyHash := utils.Base58Decode([]byte(address))
	pubKeyHash = pubKeyHash[1 : len(pubKeyHash)-4]
	UTXOs := UTXOSet.FindUTXO(pubKeyHash)

	for _, out := range UTXOs {
		balance += out.Value
	}

	fmt.Printf("Balance of '%s': %d\n", address, balance)
}
