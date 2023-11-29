# go-blockchain

## 环境
- go@1.18及以下
## 测试
- 终端一
  - export NODE_ID=3000 //windows用set替换export
  - ./go-blockchain createwallet // 将在终端上生成一个钱包地址，下面简称CENTREAL_NODE
  - ./go-blockchain createblockchain -address CENTREAL_NODE
  - cp blockchain_3000.db blockchain_genesis.db // windows用copy替换cp
- 终端二 
  - export NODE_ID=3001
  - cp blockchain_genesis.db blockchain_3001.db
  - ./go-blockchain createwallet // 生成WALLET_1
  - ./go-blockchain createwallet // 生成WALLET_2
  - ./go-blockchain createwallet // 生成WALLET_3
  - ./go-blockchain createwallet // 生成WALLET_4
- 切换终端一
  - ./go-blockchain send -from CENTREAL_NODE -to WALLET_1 -amount 10 -mine
  - ./go-blockchain send -from CENTREAL_NODE -to WALLET_2 -amount 10 -mine
  - ./go-blockchain startnode
- 切换终端二
  - ./go-blockchain startnode // 将进行终端一和终端二的同步
- 终端三
  - export NODE_ID=3002
  - cp blockchain_genesis.db blockchain_3002.db
  - ./go-blockchain createwallet // 生成MINER_WALLET
  - ./go-blockchain startnode -miner MINER_WALLET // 将与CENTREAL_NODE同步并将自身标记为矿工节点
- 切换到终端二
  - ./go-blockchain send -from WALLET_1 -to WALLET_3 -amount 1
  - ./go-blockchain send -from WALLET_2 -to WALLET_4 -amount 1 // 观察终端三，可见到有新区块生成，并且被终端一和终端三记录
  - ./go-blockchain startnode // 将与CENTREAL_NODE同步最新的区块信息
  - ./go-blockchain getbalance -address WALLET_1 // Balance of 'WALLET_1': 9
  - ./go-blockchain getbalance -address WALLET_2 // Balance of 'WALLET_2': 9
  - ./go-blockchain getbalance -address WALLET_3 // Balance of 'WALLET_3': 1
  - ./go-blockchain getbalance -address WALLET_4 // Balance of 'WALLET_4': 1
  - ./go-blockchain getbalance -address MINER_WALLET // Balance of 'MINER_WALLET': 10
