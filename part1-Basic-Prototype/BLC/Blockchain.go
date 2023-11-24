package BLC

import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"
	"math/big"
	"time"
)

// 数据库名字
const dbName = "blockchain.db"
const blockTableName = "blocks"

type Blockchain struct {
	//Blocks []*Block //Store the ordered block
	Tip []byte
	DB  *bolt.DB
}

// 迭代器
func (blockchain *Blockchain) Iterator() *BlockchainIterator {
	return &BlockchainIterator{blockchain.Tip, blockchain.DB}
}

func (blc *Blockchain) PrintChain() {
	var block *Block
	var currentHash []byte = blc.Tip
	for {
		err := blc.DB.View(func(tx *bolt.Tx) error {

			//1.表
			b := tx.Bucket([]byte(blockTableName))
			if b != nil {
				blockBytes := b.Get(currentHash)
				block = DeserializeBlock(blockBytes)

				fmt.Printf("Height: %d\n", block.Height)
				fmt.Printf("PrevBlockHash: %x\n", block.PrevBlockHash)
				fmt.Printf("Data: %s\n", block.Data)
				fmt.Printf("Timestamp: %d\n", time.Unix(block.Timestamp, 0).Format("2006-01-02 03:04:05 PM"))
				fmt.Printf("Hash: %x\n", block.Hash)
				fmt.Printf("Nonce: %d\n", block.Nonce)
			}

			return nil
		})
		if err != nil {
			log.Panic(err)
		}
		var hashInt big.Int
		hashInt.SetBytes(block.PrevBlockHash)

		if big.NewInt(0).Cmp(&hashInt) == 0 {
			break
		}
		currentHash = block.PrevBlockHash
	}

}

// 增加区块到区块链里面
func (blc *Blockchain) AddBlockToBlockchain(data string) {
	////创建新区块
	//newBlock := NewBlock(data, height, preHash)
	////往链里面添加区块
	//blc.Blocks = append(blc.Blocks, newBlock)
	err := blc.DB.Update(func(tx *bolt.Tx) error {
		//1.获取表
		b := tx.Bucket([]byte(blockTableName))
		//2.创建新区块
		if b != nil {
			blockBytes := b.Get(blc.Tip)
			block := DeserializeBlock(blockBytes)

			//3.将区块序列化并且存储在数据库中
			newBlock := NewBlock(data, block.Height+1, block.Hash)
			err := b.Put(newBlock.Hash, newBlock.Serializa())
			if err != nil {
				log.Panic(err)
			}
			//4.更新hash
			err = b.Put([]byte("l"), newBlock.Hash)
			if err != nil {
				log.Panic(err)
			}
			//5.更新blockchain的Tip
			blc.Tip = newBlock.Hash
		}
		return nil
	})
	if err != nil {
		log.Panic(err)
	}
}

// 1. 创建带有创世区块的区块链
func CreateBlockchainWithGenesisBlock() *Blockchain {
	//打开数据库
	db, err := bolt.Open(dbName, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	var blockHash []byte
	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockTableName))
		if b != nil {
			//创建数据库表
			b, err = tx.CreateBucket([]byte(blockTableName))
			if err != nil {
				log.Panic(err)
			}
		}

		if b == nil {
			//创建创世区块
			genesisBlock := CreateGenesisBlock("Genesis Data.....")
			//将创世区块存储到表中
			err := b.Put(genesisBlock.Hash, genesisBlock.Serializa())
			if err != nil {
				log.Panic(err)
			}
			//存储最新的区块的Hash
			err = b.Put([]byte("l"), genesisBlock.Hash)
			if err != nil {
				log.Panic(err)
			}
			blockHash = genesisBlock.Hash

		}
		return nil
	})

	//返回区块链对象
	//return &Blockchain{[]*Block{genesisBlock}}
	return &Blockchain{blockHash, db}
}
