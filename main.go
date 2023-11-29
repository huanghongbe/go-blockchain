package main

import (
	"fmt"
	"go-blockchain/clis"
	constants "go-blockchain/constents"
	"go-blockchain/utils"
)

func main() {
	//
	//bc := models.NewBlockchain("3000")
	//models.GetDifficulty(bc)

	cli := clis.CLI{}
	cli.Run()
}

func testNewBLock() {
	redisClient := utils.NewRedisClient()

	err := redisClient.SetIntValue(constants.BlockchainHeightKey, 1)
	if err != nil {
		return
	}

	value, err := redisClient.GetValueByKey(constants.BlockchainHeightKey)
	if err != nil {
		fmt.Println("获取键值对时发生错误:", err)
		return
	}

	fmt.Println("获取到的值为:", value)
}
func testRedisClient() {
	// 创建 Redis 客户端
	redisClient := utils.NewRedisClient()

	// 插入键值对
	err := redisClient.SetValue("myKey", "myValue")
	if err != nil {
		fmt.Println("插入键值对时发生错误:", err)
		return
	}

	fmt.Println("键值对插入成功")

	// 获取键值对
	value, err := redisClient.GetValueByKey("myKey")
	if err != nil {
		fmt.Println("获取键值对时发生错误:", err)
		return
	}

	fmt.Println("获取到的值为:", value)
}
