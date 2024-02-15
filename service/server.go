package service

import (
	"del/tools"
	"os"
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
	// 当前暂存目录
	dumpster := tools.GetCurrentUserRootDirectory() + "/.dumpster"

	// 当前暂存目录配置文件
	dumpsterConf := dumpster + "/" + "registration.rec"

	// 系统初始化
	InitRootWorkSpace(dumpsterConf)

	// 循环用户输入的路径
	for _, file := range files {
		// 将相对路径转为绝对路径
		file = tools.RelativePathToAbsolutePath(file)

		// 重新拼接暂存文件夹
		dumpster_file := dumpster + "/" + tools.GetTime()

		// 重命名/移动文件/文件夹
		err := os.Rename(file, dumpster_file)
		// 检查错误
		if err != nil {
			tools.Err("Error: Unable to delete file, check permissions and file status !", err)
		}
	}
}

// 初始化运行空间
// 用于初始化暂存目录和记述文件
func InitRootWorkSpace(dumpsterConf string) {
	// 当前用户目录
	dumpster := tools.GetCurrentUserRootDirectory() + "/.dumpster"
	if !tools.IsFileExist(dumpsterConf) {
		os.MkdirAll(dumpster, 0755)
		os.Create(dumpsterConf)
	}

}

// 获取配置文件
func GetWorkSpaceConf() {

}
