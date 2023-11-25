package BLC

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/big"
)

// 256位Hash前面至少要有16个0
const targetBit = 16

type ProofOfWork struct {
	Block  *Block   //当前要验证的区块
	target *big.Int //大数据存储
}

// 数据拼接，返回字节数组
func (pow *ProofOfWork) prepareData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.Block.PrevBlockHash,
			pow.Block.Data,
			IntToHex(pow.Block.Timestamp),
			IntToHex(int64(targetBit)),
			IntToHex(int64(nonce)),
			IntToHex(int64(pow.Block.Height)),
		},
		[]byte{},
	)

	return data
}

func (ProofOfWork *ProofOfWork) IsValid() bool {
	//1.proofOfWork.Block.Hash
	//2.proofOfWork.Target

	var hashInt big.Int
	hashInt.SetBytes(ProofOfWork.Block.Hash)

	if ProofOfWork.target.Cmp(&hashInt) == 1 {
		return true
	}
	return false
}

func (ProofOfWork *ProofOfWork) Run() ([]byte, int64) {
	//1.将Block的属性拼接成字节数组
	//2.生成Hash
	//3.判断hash有效性，如果满足条件，跳出循环
	nonce := 0
	var hashInt big.Int //存储我们新生成的Hash
	var hash [32]byte

	for {
		//准备数据
		dataBytes := ProofOfWork.prepareData(nonce)
		//生成hash
		hash = sha256.Sum256(dataBytes)
		fmt.Printf("\r%x", hash)
		//将hash存储到hashInt
		hashInt.SetBytes(hash[:])
		//判断hashInt是否小于Block里面的target
		if ProofOfWork.target.Cmp(&hashInt) == -1 {
			break
		}
		nonce = nonce + 1

	}

	return hash[:], int64(nonce)
}

// 创建新的工作量证明对象
func NewProofOfWork(block *Block) *ProofOfWork {
	//1. 创建一个初始值为1的target
	target := big.NewInt(1)
	//2.左移256 -  targetBit
	target = target.Lsh(target, 256-targetBit)
	return &ProofOfWork{block, target}

}