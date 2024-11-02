package main

import (
	"fmt"
	"os"
	"strings"

	xor "github.com/liuyuhe666/go-file-xor"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("参数错误: xor [file]")
		return
	}
	secret, ok := os.LookupEnv("XOR_SECRET")
	if !ok {
		fmt.Println("secret 不能为空")
		return
	}
	fileName := os.Args[1]
	secretBytes := []byte(secret)
	fileBytes, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Printf("读取 %s 出错: %s", fileName, err)
		return
	}
	result := xor.Xor(fileBytes, secretBytes)
	newFileName := ""
	if strings.HasSuffix(fileName, ".xor") {
		// XOR 解密
		newFileName = fileName[0 : len(fileName)-4]
	} else {
		// XOR 加密
		newFileName = fileName + ".xor"

	}
	err = os.WriteFile(newFileName, result, 0664)
	if err != nil {
		fmt.Printf("写入 %s 出错: %s", newFileName, err)
	}
}
