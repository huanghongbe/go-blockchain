package clis

import (
	"fmt"
	"go-blockchain/models"
	"go-blockchain/server"
	"log"
)

func (cli *CLI) send(from, to string, amount int, nodeID string, mineNow bool) {
	if !models.ValidateAddress(from) {
		log.Panic("ERROR: Sender address is not valid")
	}
	if !models.ValidateAddress(to) {
		log.Panic("ERROR: Recipient address is not valid")
	}

	bc := models.NewBlockchain(nodeID)
	UTXOSet := models.UTXOSet{bc}
	defer bc.DB.Close()

	wallets, err := models.NewWallets(nodeID)
	if err != nil {
		log.Panic(err)
	}
	wallet := wallets.GetWallet(from)

	tx := models.NewUTXOTransaction(&wallet, to, amount, &UTXOSet)

	if mineNow {
		cbTx := models.NewCoinbaseTX(from, "")
		txs := []*models.Transaction{cbTx, tx}

		newBlock := bc.MineBlock(txs)

		print(newBlock)
		UTXOSet.Update(newBlock)
	} else {
		server.SendTx(server.KnownNodes[0], tx)
	}

	fmt.Println("Success!")
}
