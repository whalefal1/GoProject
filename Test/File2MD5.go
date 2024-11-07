package main

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// 打开文件
	file, err := os.Open("D:\\GOLANG\\GoProject\\Test\\test.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	hasher := md5.New()
	_, err = hasher.Write(data)
	if err != nil {
		fmt.Println("Error writing to hasher:", err)
		return
	}
	hashBytes := hasher.Sum(nil)
	hashString := fmt.Sprintf("%x", hashBytes)
	fmt.Println("MD5 Hash:", hashString)
}
