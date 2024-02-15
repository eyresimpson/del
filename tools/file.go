package tools

import (
	"io/ioutil"
	"os"
)

// 判断文件夹是否存在
func IsDirExist(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

// 判断文件是否存在
func IsFileExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
		return true
	}
	return true
}

// 判断是否为文件夹
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// 判断是否为文件
func IsFile(path string) bool {
	return !IsDir(path)
}

// 写文件
func WriteFile(content string, filePath string) {
	err := ioutil.WriteFile(filePath, []byte(content), 0644)
	if err != nil {
		Err("Error writing to file: ", err)
	}
}
