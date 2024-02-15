package tools

import (
	"strconv"
	"time"
)

// 获取当前时间戳
func GetTime() string {
	return strconv.FormatInt(time.Now().UnixNano(), 10)
}
