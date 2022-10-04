// Copyright (c) Mundane of Author Createitv 2022.
// @Time        : 2022/10/3 23:23
// @Author      : Createitv
// @FileName    : time.go
// @Software    : GoLand
// @WeChat      : Navalism1
// @Description :

package time

import (
	"time"
)

// FormatTime
//  @Description: Get the year day hour min sec of given time
//  @param t
//  @return []int
func FormatTime(t time.Time) []int {
	hour, min, sec := t.Clock()
	year, month, day := t.Date()
	return []int{year, int(month), day, hour, min, sec, t.Nanosecond()}
}

// StandardFormatTime
//  @Description-es: Get the standard year-month-day hour:min:sec format
//  @Description-cn: 获取习惯性的年月日 时分秒的时间格式
//  @return string
func StandardFormatTime() string {
	tm := time.Unix(time.Now().Unix(), 0)
	return tm.Format("2006-01-02 15:04:05")
}

// Timestamp
//  @Description-es: Timestamp gets timestamp with second like 1596604455.
//  @Description-cn: 输出秒级时间戳
//  @return int64
func Timestamp() int64 {
	return time.Now().Unix()
}

// TimestampMilli
//  @Description-es: Timestamp gets timestamp with milliseconds like 1664814989514
//  @Description-cn:  获取毫秒级时间戳
//  @return int64
func TimestampMilli() int64 {
	return time.Now().UnixMilli()
}

// GetWeekday
//  @Description-es: get week day from now
//  @Description-cn: 获取今天星期几
//  @return time.Weekday
func GetWeekday() time.Weekday {
	return time.Now().Weekday()
}
