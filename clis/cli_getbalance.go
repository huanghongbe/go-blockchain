package clis

import (
	"fmt"
	"go-blockchain/models"
	"go-blockchain/utils"
	"log"
)

func (cli *CLI) getBalance(address, nodeID string) {
	if !models.ValidateAddress(address) {
		log.Panic("ERROR: Address is not valid")
	}
	bc := models.NewBlockchain(nodeID)
	UTXOSet := models.UTXOSet{bc}
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
