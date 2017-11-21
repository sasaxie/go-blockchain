package cli

import (
	"fmt"
	"github.com/sasaxie/go-blockchain/core"
	"log"
)

func (cli *CLI) createBlockchain(address, nodeID string) {
	if !core.ValidateAddress(address) {
		log.Panic("ERROR: Address is not valid")
	}
	bc := core.CreateBlockchain(address, nodeID)
	defer bc.DB.Close()

	UTXOSet := core.UTXOSet{bc}
	UTXOSet.Reindex()

	fmt.Println("Done!")
}
