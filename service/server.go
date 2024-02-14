package service

import (
	"del/tools"
	"os"
	"os/user"
	"strings"
	"time"
)

// 获取版本
func GetVersion() {
	tools.Info("TLog Version 0.0.1 Feb 5, 2024")
	tools.Info("Creator: Noah Jones")
}

// 获取帮助
func GetHelp() {
	tools.Warn("Del Version 0.0.1 Feb 5, 2024")
	tools.Info("del [-h|-v|-r|-f] file1 file2 file3 ...")
	tools.Warn("Amenity")
	tools.Info("    Enter the file you want to delete directly")
	tools.Warn("Base")
	tools.Info("    -h	 	: show help")
	tools.Info("    -v		: show version")
	tools.Info("    -f	 	: Forced deletion")
	tools.Info("    -r		: Delete a directory without prompting")
}

// 功能检查
func Run(arg string, files []string) {
	u, _ := user.Current()
	// 当前用户目录
	dumpster := u.HomeDir + "/.dumpster"
	// 当前工作目录
	path, _ := os.Executable()
	// 检查用户根目录下是否有 .dumpster 目录
	if !tools.IsDirExist(dumpster) {
		os.MkdirAll(dumpster, 0755)
	}

	dumpster += "/" + time.Now().Format("2006-01-02:11-11-11-000")
	for _, file := range files {
		// 将相对路径转为绝对路径
		if !strings.HasPrefix(file, "/") {
			file = path + "/" + file
		}
		println("file --> ", file, dumpster)
		os.MkdirAll(file, 0755)
		os.Rename(file, dumpster)
	}
}
