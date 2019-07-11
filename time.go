package goutil

import "time"

//Now 返回格式化当前时间
func Now() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

//Date 返回格式化当前日期
func Date() string {
	return time.Now().Format("2006-01-02")
}
