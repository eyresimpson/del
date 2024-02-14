package tools

import "os"

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

func IsFileExist() {

}

func IsDir() {

}

func IsFile() {

}

func delete() {

}
