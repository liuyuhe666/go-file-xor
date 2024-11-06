package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

const secretKey = "XOR_SECRET"

var rootCmd = &cobra.Command{
	Use:   "xor <filename>",
	Short: "go file xor",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Fprintln(os.Stderr, "参数错误")
			return
		}
		secret, ok := os.LookupEnv(secretKey)
		if !ok {
			fmt.Fprintln(os.Stderr, "密钥不能为空")
			return
		}
		fileName := args[0]
		secretBytes := []byte(secret)
		fileBytes, err := os.ReadFile(fileName)
		if err != nil {
			fmt.Fprintf(os.Stderr, "读取 %s 时, 发生错误: %s\n", fileName, err.Error())
			return
		}
		result := Xor(fileBytes, secretBytes)
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
			fmt.Fprintf(os.Stderr, "写入 %s 时, 发生错误: %s\n", newFileName, err.Error())
		}
	},
}

func Execute() {
	_ = rootCmd.Execute()
}
