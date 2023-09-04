package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	//模拟交易数据
	data := "hello world"
	for i := 0; i < 10000; i++ {
		hash := sha256.Sum256([]byte(data + string(i)))
		fmt.Printf("hash: %x\n", hash[:])
	}
}
