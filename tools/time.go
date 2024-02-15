package tools

import (
	"strconv"
	"time"
)

// 获取当前时间戳
func GetTime() string {
	return strconv.FormatInt(time.Now().UnixNano(), 10)
}

// 时间戳转为字符串
func TimeStampToString(timestamp int64) string {
	return time.Unix(timestamp/1000, 0).Format("2006-01-02 15:04:05")
}

// 字符串转时间戳
func StringToTimeStamp(strTime string) int64 {
	t, _ := time.Parse("2006-01-02 15:04:05", strTime)
	return t.Unix()
}
