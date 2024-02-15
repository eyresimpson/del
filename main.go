package main

import (
	"del/service"
	"del/tools"
	"os"
	"strings"
)

// 主入口函数
func main() {
	if len(os.Args) < 2 {
		tools.Err("alt-Del: Missing parameters/Incorrect number of parameters passed", nil)
		service.GetHelp()
		return
	}
	if !strings.HasPrefix(os.Args[1], "-") {
		service.Run("", os.Args[1:])
		return
	}
	switch os.Args[1] {
	case "-v":
		service.GetVersion()
	case "-h":
		service.GetHelp()
	case "-l":
		service.ListAllDeleted()
	default:
		service.Run(os.Args[1], os.Args[2:])
	}

}
