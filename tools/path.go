package tools

import "strings"

// 相对路径转绝对路径
// 仅在linux和mac os下生效，原理是判断前缀是否为 / ，如果不是则加上当前运行目录
func RelativePathToAbsolutePath(path string) string {
	if !strings.HasPrefix(path, "/") {
		path = GetCurrentDirectory() + "/" + path
	}
	return path
}
