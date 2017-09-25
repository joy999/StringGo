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

func (j JsonMap) GetJsonMap(key string) JsonMap {
	var ret JsonMap = make(JsonMap, 0)
	switch v := j[key].(type) {
	case JsonMap:
		ret = v
	case map[string]interface{}:
		ret = JsonMap(v)
	}

	return ret
}

func (j JsonMap) GetValString(key string) string {
	var ret string = ""
	switch v := j[key].(type) {
	case float64, float32:
		ret = fmt.Sprintf("%f", v)
	case string:
		ret = v
	case int, int64, int32, int16, uint64, uint16, uint32, uint8, int8:
		ret = fmt.Sprintf("%d", v)
	default:
		if v == nil {
			return ""
		}
		ret = fmt.Sprintf("%v", v)
	}

	return ret
}

func ToJsonArray(value interface{}) JsonArray {
	switch v := value.(type) {
	case JsonArray:
		return v
	case []interface{}:
		return JsonArray(v)
	case []int:
		arr := newJsonArray(len(v))
		for k, kv := range v {
			arr[k] = kv
		}
		return arr
	case []int8:
		arr := newJsonArray(len(v))
		for k, kv := range v {
			arr[k] = kv
		}
		return arr
	case []int16:
		arr := newJsonArray(len(v))
		for k, kv := range v {
			arr[k] = kv
		}
		return arr
	case []int32:
		arr := newJsonArray(len(v))
		for k, kv := range v {
			arr[k] = kv
		}
		return arr
	case []int64:
		arr := newJsonArray(len(v))
		for k, kv := range v {
			arr[k] = kv
		}
		return arr
	case []uint:
		arr := newJsonArray(len(v))
		for k, kv := range v {
			arr[k] = kv
		}
		return arr
	case []uint8:
		arr := newJsonArray(len(v))
		for k, kv := range v {
			arr[k] = kv
		}
		return arr
	case []uint16:
		arr := newJsonArray(len(v))
		for k, kv := range v {
			arr[k] = kv
		}
		return arr
	case []uint32:
		arr := newJsonArray(len(v))
		for k, kv := range v {
			arr[k] = kv
		}
		return arr
	case []uint64:
		arr := newJsonArray(len(v))
		for k, kv := range v {
			arr[k] = kv
		}
		return arr
	case []string:
		arr := newJsonArray(len(v))
		for k, kv := range v {
			arr[k] = kv
		}
		return arr
	case []float32:
		arr := newJsonArray(len(v))
		for k, kv := range v {
			arr[k] = kv
		}
		return arr
	case []float64:
		arr := newJsonArray(len(v))
		for k, kv := range v {
			arr[k] = kv
		}
		return arr
	default:
		arr := NewJsonArray()
		return arr
	}
}

func (this JsonMap) GetValArray(key string) JsonArray {
	return ToJsonArray(key)
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

func NewJsonArray() JsonArray {
	return newJsonArray(0)
}

func newJsonArray(size int) JsonArray {
	return make(JsonArray, size)
}

func (j JsonArray) JsonEncode() (JsonString, error) {
	s, err := json.Marshal(j)
	return JsonString(s), err
}

func (j JsonArray) Append(v interface{}) {
	j = append(j, v)
}

func (j JsonArray) CopyFrom(arr JsonArray) {
	copy(j, JsonArray(arr))
}

func (j JsonArray) Resize(n int) {
	if n <= len(j) {
		j = j[:n]
	} else {
		t := j
		j = newJsonArray(n)
		copy(j, t)
	}
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
