package utils

import (
	"errors"
	"fmt"
	"os"
)

func PathExist(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, errors.New("文件已存在")
	}
	return false, err
}

func GetHostname() (string, error) {
	hostnameFile := "/etc/hostname"
	hostname, err := os.ReadFile(hostnameFile)
	if err != nil {
		fmt.Printf("Error reading hostname file: %s\n", err)
		return "unkown", err
	}
	return string(hostname), nil
}
