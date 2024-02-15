package tools

import (
	"fmt"
)

// 日志模块

func Info(text any) {
	fmt.Printf("\033[0;34m%s\033[0m\n", text)

}

func Warn(text any) {
	fmt.Printf("\033[0;33m%s\033[0m\n", text)
}

func Success(text any) {
	fmt.Printf("\033[0;32m%s\033[0m\n", text)
}

func Err(text any, err error) {
	fmt.Printf("\033[1;31m%s\033[0m\n", text)
	if err != nil {
		fmt.Printf("\033[1;31m%s\033[0m\n", "\t"+err.Error())
	}
}
