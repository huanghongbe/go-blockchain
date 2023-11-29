package models

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"github.com/boltdb/bolt"
	constants "go-blockchain/constents"
	"go-blockchain/utils"
	"log"
	"math"
	"math/big"
)

var (
	maxNonce = math.MaxInt64
)

// ProofOfWork represents a proof-of-work
type ProofOfWork struct {
	block  *Block
	target *big.Int
}

func GetDifficulty(bc *Blockchain) int {
	if bc == nil {
		return constants.DefaultDifficulty
	}
	var lastHash []byte
	var lastHeight int
	var lastDifficulty int
	err := bc.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		lastHash = b.Get([]byte("l"))

		blockData := b.Get(lastHash)
		block := DeserializeBlock(blockData)

		lastHeight = block.Height
		lastDifficulty = block.Difficulty
		return nil
	})
	if err != nil {
		log.Panic(err)
	}
	if lastHeight < 10 {
		return constants.DefaultDifficulty
	}
	var currentDifficulty = constants.DefaultDifficulty

	duration := calculateDuration(bc)
	avgBlockTime := duration / 10
	fmt.Println("前十个块用时", duration)
	if avgBlockTime > 600*1.5 {
		currentDifficulty = lastDifficulty - 1
	} else if avgBlockTime < 600*0.75 {
		currentDifficulty = lastDifficulty + 1
	}
	fmt.Println("Difficulty调整", currentDifficulty)
	return currentDifficulty
}

func calculateDuration(bc *Blockchain) int64 {
	iterator := bc.Iterator()
	blockCount := 0
	var firstBlock *Block
	var lastBlock *Block
	for block := iterator.Next(); block != nil; block = iterator.Next() {
		if firstBlock == nil {
			firstBlock = block
		}
		lastBlock = block
		blockCount++
		if blockCount > 10 {
			break
		}
	}
	duration := firstBlock.Timestamp - lastBlock.Timestamp
	return duration
}

// NewProofOfWork builds and returns a ProofOfWork
func NewProofOfWork(b *Block, bc *Blockchain) *ProofOfWork {
	target := big.NewInt(1)
	//todo : 动态获取difficulty , 并设置给block
	difficulty := GetDifficulty(bc)
	b.Difficulty = difficulty
	target.Lsh(target, uint(256-difficulty))

	pow := &ProofOfWork{b, target}

	return pow
}

func (pow *ProofOfWork) prepareData(nonce int) []byte {

	//pow.block.Difficulty = getDifficulty(nodeID)
	data := bytes.Join(
		[][]byte{
			pow.block.PrevBlockHash,
			pow.block.HashTransactions(),
			utils.IntToHex(pow.block.Timestamp),
			utils.IntToHex(int64(pow.block.Difficulty)),
			utils.IntToHex(int64(nonce)),
		},
		[]byte{},
	)

	return data
}

// Run performs a proof-of-work
func (pow *ProofOfWork) Run() (int, []byte) {
	var hashInt big.Int
	var hash [32]byte
	nonce := 0

	fmt.Printf("Mining a new block")
	for nonce < maxNonce {

		data := pow.prepareData(nonce)

		hash = sha256.Sum256(data)
		if math.Remainder(float64(nonce), 100000) == 0 {
			fmt.Printf("\r%x", hash)
		}
		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(pow.target) == -1 {
			break
		} else {
			nonce++
		}
	}
	fmt.Print("\n\n")

	return nonce, hash[:]
}

// Validate validates block's PoW
func (pow *ProofOfWork) Validate() bool {
	var hashInt big.Int

	data := pow.prepareData(pow.block.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])

	isValid := hashInt.Cmp(pow.target) == -1

	return isValid
}
