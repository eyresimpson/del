package service

import (
	"del/tools"
	"os"
	"strconv"
)

// 获取版本
func GetVersion() {
	tools.Info("TLog Version 0.1.0 Feb 15, 2024")
	tools.Info("Creator: Noah Jones")
}

// 获取帮助
func GetHelp() {
	tools.Warn("Del Version 0.1.0 Feb 15, 2024")
	tools.Info("del [-h|-v|-r|-f] file1 file2 file3 ...")
	tools.Warn("Amenity")
	tools.Info("    Enter the file you want to delete directly")
	tools.Warn("Base")
	tools.Info("    -h	 	: show help")
	tools.Info("    -v		: show version")
	tools.Info("    -f	 	: Forced deletion")
	// 列出所有已删除
	tools.Info("    -l		: Lists all deleted files")
	// 恢复指定文件
	tools.Info("    -r		: Recover specified file, use id")
}

// 功能检查
func Run(arg string, files []string) {
	// 当前暂存目录
	dumpster := tools.GetCurrentUserRootDirectory() + "/.dumpster"

	// 当前暂存目录配置文件

	// 系统初始化
	InitRootWorkSpace(GetConfigPath())

	// 循环用户输入的路径
	for index, file := range files {
		// 将相对路径转为绝对路径
		file = tools.RelativePathToAbsolutePath(file)

		timeStamp := tools.GetTime()
		// 重新拼接暂存文件夹
		dumpster_file := dumpster + "/" + timeStamp + "_" + strconv.Itoa(index)

		// 将文件信息写入配置文件
		// 原始文件位置
		tools.WriteJsonArray(GetConfigPath(), map[string]string{"id": timeStamp + "_" + strconv.Itoa(index), "rawPath": file, "dumpPath": dumpster_file, "isDir": strconv.FormatBool(tools.IsDir(file))})

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
		tools.WriteFile("[]", dumpsterConf)
	}

}

// 获取配置文件
func GetWorkSpaceConf() {

}

func ListAllDeleted() {

	// 读取配置文件
	data := tools.ReadJsonArray(GetConfigPath())

	if len(data) == 0 {
		tools.Info("No deleted file found !")
	}

	for _, val := range data {
		if val["isDir"] == "true" {
			tools.Info("Dir : " + val["id"])
		} else {
			tools.Info("File : " + val["id"])
		}
		tools.Info("\tRaw  Path: " + val["rawPath"])
		tools.Info("\tDump Path: " + val["dumpPath"])
	}

}

func GetConfigPath() string {
	return tools.GetCurrentUserRootDirectory() + "/.dumpster/registration.json"
}
