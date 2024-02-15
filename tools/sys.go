package tools

import (
	"bytes"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
)

func RunExe(exe string, args ...string) error {
	command := exec.Command(exe, args...)
	// 此处同步执行，为了执行完成后加个提示
	err := command.Run()
	return err
}

// 模拟系统的命令行运行，运行日志同步写入指定目录
func Run(command string, args ...string) string {
	cmd := exec.Command(command, args...)
	var out, err bytes.Buffer
	cmd.Stderr = &err
	cmd.Stdout = &out
	cmd.Run()
	errStr := string(err.String())
	outStr := string(out.String())
	if errStr != "" {
		return errStr
	} else {
		return outStr
	}
}

// 获取当前运行目录
func GetCurrentDirectory() string {
	rawPath, _ := os.Executable()
	return filepath.Dir(rawPath)
}

// 获取用户根目录
func GetCurrentUserRootDirectory() string {
	u, _ := user.Current()
	return u.HomeDir
}
