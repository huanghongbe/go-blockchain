package main

import (
	"github.com/boltdb/bolt"
	"log"
	"part1_Basic_Prototype/BLC"
	_ "part1_Basic_Prototype/BLC"
)

func main() {

	//创世区块
	blockchain := BLC.CreateBlockchainWithGenesisBlock()
	defer blockchain.DB.Close()
	//新区块
	blockchain.AddBlockToBlockchain("Send 100RMB To zhangqiang")
	blockchain.AddBlockToBlockchain("Send 200RMB To changjingkong")
	blockchain.AddBlockToBlockchain("Send 300RMB To juncheng")
	blockchain.AddBlockToBlockchain("Send 500RMB To haolin")

	blockchain.PrintChain()

	//block := BLC.NewBlock("tEST", 1, []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
	//fmt.Printf("%d", block.Nonce)
	//fmt.Printf("%x", block.Hash)

	//bytes := block.Serializa()
	//fmt.Println(bytes)
	//block = BLC.DeserializeBlock(bytes)
	//fmt.Printf("%d", block.Nonce)
	//fmt.Printf("%x", block.Hash)

	//打开数据库
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

}
