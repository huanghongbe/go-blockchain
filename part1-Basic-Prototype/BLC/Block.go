package BLC

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
	"log"
	"strconv"
	"time"
)

type Block struct {
	//1.Height of Block
	Height int64
	//2.上一个区块的Hash
	PrevBlockHash []byte
	//3.Transcation Data
	Data []byte
	//4.TimeStamp时间戳
	Timestamp int64
	//5.Hash
	Hash []byte
	//6. Nonce新鲜值
	Nonce int64
}

// 将区块链序列化成字节数组
func (block *Block) Serializa() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(block)
	if err != nil {
		log.Panic(err)
	}

	return result.Bytes()

}

// 反序列化
func DeserializeBlock(blockBytes []byte) *Block {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(blockBytes))
	err := decoder.Decode(&block)
	if err != nil {
		log.Panic(err)
	}

	return &block
}

func (block *Block) SetHash() {
	//1.Height
	heightBytes := IntToHex(block.Height)
	fmt.Println(heightBytes)
	//2.Transfer the TimeStamp to []byte
	timeString := strconv.FormatInt(block.Timestamp, 2)

	timeBytes := []byte(timeString)
	fmt.Println(timeBytes)
	//3.拼接所有属性
	blockBytes := bytes.Join([][]byte{heightBytes, block.PrevBlockHash, block.Data, timeBytes, block.Hash}, []byte{})
	//4.生成Hash
	hash := sha256.Sum256(blockBytes)

	block.Hash = hash[:]
}

// 1. Create a new block
func NewBlock(data string, height int64, preBlockHash []byte) *Block {
	//Create the Block
	block := &Block{height, preBlockHash, []byte(data), time.Now().Unix(), nil, 0}
	//调用工作量证明的方法并且返回有效的Hash和Nonce
	pow := NewProofOfWork(block)

	hash, nonce := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce
	return block
}

// 2. 生成创世区块
func CreateGenesisBlock(data string) *Block {

	return NewBlock(data, 1, []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
}
