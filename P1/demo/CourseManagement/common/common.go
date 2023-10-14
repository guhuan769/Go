package common

import "time"

type Gender int

// 常量
const (
	//女
	Female Gender = 0
	//男
	Male Gender = 1
	//第三性别
	Third Gender = 2
)

type TimeLayout string

const Date TimeLayout = "2006-01-02"
const DateTime TimeLayout = "2006-01-02 15:04:05"
const DateTimeZone TimeLayout = "2006-01-02 15:04:05Z0700"

// ... 可选 可以有可以没有 字符串转时间
func StrToTime(str string, layout ...TimeLayout) (time.Time, error) {
	format := Date
	if len(layout) > 0 {
		format = layout[0]
	}
	return time.Parse(string(format), str)
}

// 时间转字符串
func TimeToStr(t time.Time, layout ...TimeLayout) string {
	format := Date
	if len(layout) > 0 {
		format = layout[0]
	}
	return t.Format(string(format))
}
