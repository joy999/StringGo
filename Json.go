package utils

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type JsonMap map[string]interface{}

type JsonResult JsonMap

func NewJsonMap() JsonMap {
	m := make(JsonMap, 0)
	return m
}

func (j JsonMap) JsonEncode() (JsonString, error) {
	s, err := json.Marshal(j)
	return JsonString(s), err
}

func (j JsonMap) GetValString(key string) string {
	var ret string
	switch v := j[key].(type) {
	case float64, float32:
		ret = fmt.Sprintf("%f", v)
	case string:
		ret = v
	case int, int64, int32, int16, uint64, uint16, uint32, uint8, int8:
		ret = fmt.Sprintf("%d", v)
	default:
		ret = fmt.Sprintf("%v", v)
	}

	return ret
}

func (j JsonMap) GetValInt64(key string) int64 {
	var ret int64
	if val, ok := j[key]; ok {
		switch v := val.(type) {
		case float64:
			ret = int64(v)
		case float32:
			ret = int64(v)
		case string:
			i, _ := strconv.Atoi(v)
			ret = int64(i)
		case int:
			ret = int64(v)
		case int64:
			ret = int64(v)
		case int32:
			ret = int64(v)
		case int16:
			ret = int64(v)
		case uint64:
			ret = int64(v)
		default:
			ret = 0
		}
	} else {
		ret = 0
	}

	return ret
}

type JsonInt int

func (j JsonInt) JsonEncode() (JsonString, error) {
	s, err := json.Marshal(j)

	return JsonString(s), err
}

func (j JsonInt) ToInt() int {
	return int(j)
}

type JsonFloat float64

func (j JsonFloat) JsonEncode() (JsonString, error) {
	s, err := json.Marshal(j)
	return JsonString(s), err
}

func (j JsonFloat) ToFloat64() float64 {
	return float64(j)
}

type JsonArray []interface{}

func (j JsonArray) JsonEncode() (JsonString, error) {
	s, err := json.Marshal(j)
	return JsonString(s), err
}

func (j JsonArray) Append(v interface{}) {
	j = append(j, v)
}

type JsonString String

func (j JsonString) JsonEncode() (JsonString, error) {
	s, err := json.Marshal(j)
	return JsonString(s), err
}

func (j JsonString) JsonDecode(t interface{}) error {
	err := json.Unmarshal([]byte(j), t)
	return err
}

func (j JsonString) ToString() String {
	return String(j)
}

func (j JsonString) ToNativeString() string {
	return string(j)
}
