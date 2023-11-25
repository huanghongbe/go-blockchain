package clis

import (
	"fmt"
	"go-blockchain/models"
	"log"
)

func (cli *CLI) createBlockchain(address, nodeID string) {
	if !models.ValidateAddress(address) {
		log.Panic("ERROR: Address is not valid")
	}
	bc := models.CreateBlockchain(address, nodeID)
	defer bc.DB.Close()

	UTXOSet := models.UTXOSet{bc}
	UTXOSet.Reindex()

	fmt.Println("Done!")
}
