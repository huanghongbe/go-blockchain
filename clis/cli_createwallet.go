package clis

import (
	"fmt"
	"go-blockchain/models"
)

func (cli *CLI) createWallet(nodeID string) {
	wallets, _ := models.NewWallets(nodeID)
	address := wallets.CreateWallet()
	wallets.SaveToFile(nodeID)

	fmt.Printf("Your new address: %s\n", address)
}
