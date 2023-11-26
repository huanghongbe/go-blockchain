package clis

import (
	"fmt"
	constants "go-blockchain/constents"
	"go-blockchain/models"
	"go-blockchain/utils"
	"log"
)

func (cli *CLI) getChainHeight(nodeID string) {
	bc := models.NewBlockchain(nodeID)
	defer bc.DB.Close()

	redis := utils.NewRedisClient() // 在这里处理获取到的链高度
	// 从 Redis 中获取区块高度
	height, err := redis.GetIntValueByKey(constants.BlockchainHeightKey)
	if err != nil {
		// 如果 Redis 中不存在区块高度，则从 BoltDB 中读取并写入 Redis
		height, err = models.GetBlockchainHeightFromBoltDB(bc.DB)
		if err != nil {
			log.Panic(err)
		}
		err = redis.SetIntValue(constants.BlockchainHeightKey, height)
		if err != nil {
			log.Panic(err)
		}
	}
	// 打印链高度信息
	fmt.Printf("Blockchain Height: %d\n", height)
}
