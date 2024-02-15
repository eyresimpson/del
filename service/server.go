package service

import (
	"del/tools"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
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
func Run(opt string, args []string) {

	// 判断参数
	switch opt {
	case "-r":
		Recover(args)
	case "-f":
		ReallyRemove(args)
	default:
		Remove(args)
	}
}

// 初始化运行空间
// 用于初始化暂存目录和记述文件
func InitRootWorkSpace(dumpsterConf string) {
	// 判断配置文件是否存在
	if !tools.IsFileExist(dumpsterConf) {
		// 获取回收站目录
		dumpster := tools.GetCurrentUserRootDirectory() + "/.dumpster"
		os.MkdirAll(dumpster, 0755)
		os.Create(dumpsterConf)
		tools.WriteFile("[]", dumpsterConf)
	}

}

func ListAllDeleted() {

	// 读取配置文件
	data := tools.ReadJsonArray(GetConfigPath())

	if len(data) == 0 {
		tools.Info("No deleted file found !")
	}

	for _, val := range data {
		if val["isDir"] == "true" {
			tools.Info("Dir [ " + val["name"] + " ]: " + val["id"])
		} else {
			tools.Info("File [ " + val["name"] + " ]: " + val["id"])
		}
		tools.Info("\tRaw  Path: " + val["rawPath"])
		tools.Info("\tDump Path: " + val["dumpPath"])
	}

}

func CleanAllDeleted() {
	var userInput string

	tools.Err("Are you sure you want to delete all the recycle bin files?", nil)
	tools.Err("This will destroy all stored recycle bin files and cannot be recovered", nil)
	tools.Warn("Type 'Y', 'y' or 'yes' to confirm, Enter anything else to cancel the operation:")
	_, err := fmt.Scanln(&userInput)
	if err != nil {
		fmt.Println("Error Load Input:", err)
		return
	}

	// 将用户输入转换为小写，并去除空格
	userInput = strings.ToLower(strings.TrimSpace(userInput))

	// 检查用户输入是否符合条件
	if userInput == "y" || userInput == "yes" {
		os.RemoveAll(tools.GetCurrentUserRootDirectory() + "/.dumpster")
		tools.Success("Operation submitted successfully!")
	} else {
		tools.Warn("Operation cancelled!")
		return
	}
	// 直接删除了临时目录

}

func GetConfigPath() string {
	return tools.GetCurrentUserRootDirectory() + "/.dumpster/registration.json"
}

func Recover(args []string) {
	// 读取配置文件
	data := tools.ReadJsonArray(GetConfigPath())

	if len(data) == 0 {
		tools.Info("No deleted file found !")
	}

	// 将指定的 id 存储在 map 中以便快速查找
	specifiedIDMap := make(map[string]bool)
	for _, id := range args {
		specifiedIDMap[id] = true
	}

	// 循环其中的所有行
	for _, val := range data {
		if specifiedIDMap[val["id"]] {
			// 重命名/移动文件/文件夹
			err := os.Rename(val["dumpPath"], val["rawPath"])

			// 检查错误
			if err != nil {
				tools.Err("Error: Unable to delete file, check permissions and file status !", err)
				return
			}
			tools.RemoveObjectFromJsonArray(GetConfigPath(), val["id"])
		}
	}

}

func Remove(args []string) {
	// 当前暂存目录
	dumpster := tools.GetCurrentUserRootDirectory() + "/.dumpster"

	// 循环用户输入的路径
	for index, file := range args {
		// 将相对路径转为绝对路径
		file = tools.RelativePathToAbsolutePath(file)

		timeStamp := tools.GetTime()
		// 重新拼接暂存文件夹
		dumpster_file := dumpster + "/" + timeStamp + "_" + strconv.Itoa(index)

		// 重命名/移动文件/文件夹
		err := os.Rename(file, dumpster_file)

		// 检查错误
		if err != nil {
			tools.Err("Error: Unable to delete file, check permissions and file status !", err)
			return
		}

		// 将文件信息写入配置文件
		// 原始文件位置
		tools.InsertObjectToJsonArray(GetConfigPath(), map[string]string{"id": timeStamp + "_" + strconv.Itoa(index), "name": filepath.Base(file), "rawPath": file, "dumpPath": dumpster_file, "isDir": strconv.FormatBool(tools.IsDir(file))})

	}
}

// 直接删除，不走del（相当于系统删除）
func ReallyRemove(args []string) {
	var userInput string

	tools.Err("Are you sure you want to delete files/dirs?", nil)
	tools.Err("This will call the system directly to delete!", nil)
	tools.Warn("Type 'Y', 'y' or 'yes' to confirm, Enter anything else to cancel the operation:")
	_, err := fmt.Scanln(&userInput)
	if err != nil {
		fmt.Println("Error Load Input:", err)
		return
	}

	// 将用户输入转换为小写，并去除空格
	userInput = strings.ToLower(strings.TrimSpace(userInput))

	// 检查用户输入是否符合条件
	if userInput == "y" || userInput == "yes" {
		// 删除指定文件
		for _, file := range args {
			os.RemoveAll(file)
		}
		tools.Success("Delete operation successfully!")

	} else {
		tools.Warn("Operation cancelled!")
		return
	}

}
