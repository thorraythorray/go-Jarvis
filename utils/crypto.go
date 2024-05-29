package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

func GetFileMD5(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return "", err
	}
	defer file.Close()

	// 创建一个新的md5哈希
	hash := md5.New()

	// 将文件内容复制到哈希
	if _, err := io.Copy(hash, file); err != nil {
		fmt.Println("Error hashing file:", err)
		return "", err
	}

	// 将哈希值转换为十六进制字符串
	hashInBytes := hash.Sum(nil) // 注意：Sum的参数是nil，表示将哈希值复制到一个新的字节切片中
	hashStr := hex.EncodeToString(hashInBytes)
	return hashStr, nil
}
