package utils

import (
	"strconv"
	"time"
)

type Time_t int64

type MicroTime_t int64

/*
  获取Unix时间戳
*/
func Time() Time_t {
	return Time_t(time.Now().Unix())
}

/**
* 获取Unix微秒级的时间戳
 */
func MicroTime() MicroTime_t {
	return MicroTime_t(time.Now().UnixNano())
}

func (t Time_t) ToString() string {
	return strconv.FormatInt(int64(t), 10)
}

func (t MicroTime_t) ToString() string {
	return strconv.FormatInt(int64(t), 10)
}

func (t Time_t) ToInt64() int64 {
	return int64(t)
}

func (t MicroTime_t) ToInt64() int64 {
	return int64(t)
}

func (t Time_t) Format(f string) string {
	fs := String(f)
	//t.Format("2006-01-02 15:04:05 Y-m-d H:i:s
	fs = fs.MatchReplace("Y", "2006")
	fs = fs.MatchReplace("m", "01")
	fs = fs.MatchReplace("d", "02")
	fs = fs.MatchReplace("H", "15")
	fs = fs.MatchReplace("i", "04")
	fs = fs.MatchReplace("s", "05")

	f = fs.ToString()

	return time.Unix(t.ToInt64(), 0).Format(f)
}

func Sleep(seconds int) {
	_s := time.Duration(seconds)
	_s *= 1e9
	time.Sleep(_s)
}

func SleepF(seconds float64) {
	_s := time.Duration(seconds * 1e9)
	time.Sleep(_s)
}
