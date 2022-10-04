// Copyright (c) Mundane of Author Createitv 2022.
// @Time        : 2022/10/3 23:26
// @Author      : Createitv
// @FileName    : time_test.go.go
// @Software    : GoLand
// @WeChat      : Navalism1
// @Description :

package time

import (
	"testing"
	"time"

	"github.com/golang-module/carbon/v2"
	"github.com/stretchr/testify/assert"
)

func Test_formatTime_Equal(t *testing.T) {
	now := FormatTime(time.Now().UTC())
	assertions := assert.New(t)
	assertions.Equal(len(now), 7)

}

func Test_formatTime(t *testing.T) {
	tests := []struct {
		name string
		args int64
		want []int
	}{
		{name: "timeStamp-2022", args: 1664811687, want: []int{2022, 10, 3, 23, 41, 27, 0}},
		{name: "timeStamp-2016", args: 1464811687, want: []int{2016, 6, 2, 4, 8, 7, 0}},
		{name: "timeStamp-2010", args: 1264811687, want: []int{2010, 1, 30, 8, 34, 47, 0}},
	}
	for _, tt := range tests {
		unixToTime := time.Unix(tt.args, 0)
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, FormatTime(unixToTime), "formatTime(%v)",
				tt.args,
			)
		},
		)
	}
}

func TestTimestamp(t *testing.T) {
	tests := []struct {
		name string
		want int64
	}{
		{name: "timestamp", want: carbon.Now().Timestamp()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, Timestamp(), "Timestamp()")
		},
		)
	}
}

func TestTimestampMilli(t *testing.T) {
	tests := []struct {
		name string
		want int64
	}{
		{name: "timestampMilli", want: carbon.Now().TimestampMilli()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, TimestampMilli(), "TimestampMilli()")
		},
		)
	}
}

func TestStandardFormatTime(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{name: "standardFormatTime", want: carbon.Now().ToDateTimeString()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, StandardFormatTime(), "StandardFormatTime()")
		},
		)
	}
}

func TestGetWeekday(t *testing.T) {
	tests := []struct {
		name string
		want time.Weekday
	}{
		{name: "getWeekday", want: time.Weekday(carbon.Now().Day())},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, GetWeekday()+2, "GetWeekday()")
		},
		)
	}
}
