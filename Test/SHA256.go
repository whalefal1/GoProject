package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func main() {
	// 待加密的字符串
	input := "Hello, World!"
	// 创建一个新的SHA-256哈希计算器
	hasher := sha256.New()
	// 写入待加密的数据
	_, err := hasher.Write([]byte(input))
	if err != nil {
		fmt.Println("Error while writing to hasher:", err)
		return
	}
	// 计算最终的哈希值
	hashBytes := hasher.Sum(nil)
	// 将哈希值转换为16进制字符串
	hashString := hex.EncodeToString(hashBytes)
	// 打印结果
	fmt.Println("SHA-256 Hash:", hashString)
}
