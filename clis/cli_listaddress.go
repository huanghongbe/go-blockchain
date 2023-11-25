package clis

import (
	"fmt"
	"go-blockchain/models"
	"log"
)

func (cli *CLI) listAddresses(nodeID string) {
	wallets, err := models.NewWallets(nodeID)
	if err != nil {
		log.Panic(err)
	}
	addresses := wallets.GetAddresses()

	for _, address := range addresses {
		fmt.Println(address)
	}
}
