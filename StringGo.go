package utils

import (
	"github.com/joy999/mahonia"
	"io"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type String string

type StringArray []String

//新创建一个空的String对象指针
func NewString() *String {
	str := new(String)
	*str = ""
	return str
}

func NewStringArray() StringArray {
	o := make([]String, 0)
	return o
}

//获取长度
func (str String) GetLength() int {
	return len(str)
}

//获取Reader
func (str String) NewReader() io.Reader {
	var r io.Reader = strings.NewReader(string(str))
	return r
}

//获取writer
func (str *String) NewWriter() io.Writer {
	var w io.Writer = str
	return w
}

func (str *String) Write(p []byte) (n int, err error) {
	*str = (*str) + String(p)

	return len(p), nil
}

//转为字节数组
func (str String) ToBytes() []byte {
	return []byte(str)
}

//转为string
func (str String) ToString() string {
	return string(str)
}

/**
检查是否符合正则
*/
func (str String) Match(pattern string) bool {
	m, err := regexp.MatchString(pattern, string(str))
	if err != nil {
		log.Fatal(err.Error())
	}
	return m
}

//获取符合正则的字符串
func (str String) MatchFind(pattern string) []String {
	reg, err := regexp.Compile(pattern)
	if err != nil {
		log.Fatal(err.Error())
	}

	s := reg.FindStringSubmatch(string(str))
	if s == nil {
		return nil
	}
	ret := make([]String, 0)
	for _, v := range s {
		ret = append(ret, String(v))
	}
	return ret
}

//获取所有符合正则的字符串
func (str String) MatchAllFind(pattern string) [][]String {
	reg, err := regexp.Compile(pattern)

	if err != nil {
		log.Fatal(err.Error())
	}
	s := reg.FindAllStringSubmatch(string(str), -1)

	if s == nil {
		return nil
	}

	ret := make([][]String, 0)
	for _, vv := range s {
		_vv := make([]String, 0)
		for _, v := range vv {
			_vv = append(_vv, String(v))
		}
		ret = append(ret, _vv)
	}
	return ret
}

//利用正则进行替换
func (str String) MatchReplace(pattern string, to string) String {
	reg, err := regexp.Compile(pattern)
	if err != nil {
		log.Fatal(err.Error())
	}

	return String(reg.ReplaceAllString(string(str), to))
}

//由GBK编码转为UTF8
func (str String) GBKToUTF8() String {
	_c := doGBKToUTF8(string(str))
	return String(_c)
}

//由UTF8转为GBK编码
func (str String) UTF8ToGBK() String {
	_c := doUTF8ToGBK(string(str))
	return String(_c)
}

func (str String) ToInt() int {
	i, err := strconv.Atoi(string(str))
	if err != nil {
		i = 0
	}
	return i
}

func (str String) ToInt64() int64 {
	i, err := strconv.ParseInt(str.ToString(), 10, 64)
	if err != nil {
		i = 0
	}
	return i
}

func (str String) Explode(sep string) StringArray {
	s := strings.Split(string(str), sep)
	ret := make([]String, 0)
	for _, v := range s {
		ret = append(ret, String(v))
	}
	return StringArray(ret)
}

func (this String) IsSame(s string) bool {
	return this.ToString() == s
}

func (this String) TrimSpace() String {
	return String(strings.TrimSpace(this.ToString()))
}

func (this String) ToUnixLocalTimeStamp(format string) int64 {
	fs := String(format)
	fs = fs.MatchReplace("Y", "2006")
	fs = fs.MatchReplace("m", "01")
	fs = fs.MatchReplace("d", "02")
	fs = fs.MatchReplace("H", "15")
	fs = fs.MatchReplace("i", "04")
	fs = fs.MatchReplace("s", "05")
	stamp, _ := time.ParseInLocation(fs.ToString(), this.ToString(), time.Local)
	return stamp.Unix()
}

func (this StringArray) Have(item string) bool {
	for _, v := range this {
		if v.IsSame(item) {
			return true
		}
	}

	return false
}

func (this StringArray) Implode(sep string) String {
	arr := this.ToNativeStringArray()

	ret := strings.Join(arr, sep)

	return String(ret)
}

//转换为原生字符串数组 []string
func (this StringArray) ToNativeStringArray() []string {
	arr := make([]string, 0)
	for _, v := range this {
		arr = append(arr, v.ToString())
	}
	return arr
}

func (this StringArray) Erase(key int) StringArray {
	n := len(this)
	if n == 0 || key < 0 || key >= n {
		return this
	}

	//ret := make(StringArray, 0)
	var ret StringArray
	if key == 0 {
		if n == 1 {

		} else {
			//ret = append(ret, this[1:]...)
			ret = append(this[1:])
		}

	} else if key == n-1 {
		//ret = append(ret, this[:n-1]...)
		ret = append(this[:n-1])
	} else {
		//ret = append(ret, this[:key]...)
		//ret = append(ret, this[key+1:]...)
		ret = append(this[:key], this[key+1:]...)
	}

	return ret
}

func (this StringArray) Push(item string) {
	this = append(this, String(item))
}

func doGBKToUTF8(inputStr string) string {
	return doEncodingConvert(inputStr, "gbk", "utf-8")
}

func doUTF8ToGBK(inputStr string) string {
	return doEncodingConvert(inputStr, "utf-8", "gbk")
}

func doEncodingConvert(inputStr string, fromCode string, toCode string) string {
	outStr := String("")

	outPointer := &outStr
	reader := String(inputStr).NewReader()
	var r io.Reader = reader
	var w io.Writer = outPointer

	if fromCode != "utf-8" {
		decode := mahonia.NewDecoder(fromCode)
		if decode == nil {
			log.Fatalf("Could not create decoder for %s", fromCode)

		}
		r = decode.NewReader(r)
	}

	if toCode != "utf-8" {
		encode := mahonia.NewEncoder(toCode)
		if encode == nil {
			log.Fatalf("Could not create encoder for %s", toCode)
		}
		w = encode.NewWriter(w)
	}
	io.Copy(w, r)

	return string(outStr)
}
