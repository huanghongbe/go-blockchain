package main

import (
	"github.com/boltdb/bolt"
	"log"
)

func main() {
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//创建表
	err = db.Update(func(tx *bolt.Tx) error {
		//创建BlockBucket表
		b := tx.Bucket([]byte("BlockBucket"))

		if b != nil {
			err := b.Put([]byte("l"), []byte("Send 100 BTC TO Someone"))
			if err != nil {
				log.Panic("数据存储失败")
			}
		}
		return nil
	})
	//更新失败
	if err != nil {
		log.Panic(err)
	}

}
