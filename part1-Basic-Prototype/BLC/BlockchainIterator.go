package BLC

import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"
	"math/big"
	"time"
)

type BlockchainIterator struct {
	CurrentHash []byte
	DB          *bolt.DB
}

func (blockchainIterator *BlockchainIterator) Next() *Block {
	var block *Block
	err := blockchainIterator.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockTableName))

		if b != nil {
			currentBlockBytes := b.Get(blockchainIterator.CurrentHash)
			block := DeserializeBlock(currentBlockBytes)
			blockchainIterator.CurrentHash = block.PrevBlockHash
		}
		return nil
	})

	if err != nil {
		log.Panic(err)
	}
	return block
}

// 遍历输出所有区块的信息
func (blc *Blockchain) Printchain() {
	blochchainIterator := blc.Iterator()

	for {
		block := blochchainIterator.Next()
		fmt.Printf("Height: %d\n", block.Height)
		fmt.Printf("PrevBlockHash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Timestamp: %d\n", time.Unix(block.Timestamp, 0).Format("2006-01-02 03:04:05 PM"))
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Printf("Nonce: %d\n", block.Nonce)

		fmt.Println()
		var hashInt big.Int
		hashInt.SetBytes(block.PrevBlockHash)

		if big.NewInt(0).Cmp(&hashInt) == 0 {
			break
		}
	}

}
