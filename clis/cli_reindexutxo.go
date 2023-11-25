package clis

import (
	"fmt"
	"go-blockchain/models"
)

func (cli *CLI) reindexUTXO(nodeID string) {
	bc := models.NewBlockchain(nodeID)
	UTXOSet := models.UTXOSet{bc}
	UTXOSet.Reindex()

	count := UTXOSet.CountTransactions()
	fmt.Printf("Done! There are %d transactions in the UTXO set.\n", count)
}
