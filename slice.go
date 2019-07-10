package goutil

import (
	"reflect"
)

//InSlice 判断是否存在slice 中
func InSlice(n, slice interface{}) (ok bool) {
	var s []interface{}
	if s, ok = slice.([]interface{}); !ok {
		return false
	}
	if len(s) == 0 {
		return false
	}
	if reflect.TypeOf(n) != reflect.TypeOf(s[0]) {
		return false
	}
	for _, v := range s {
		if n == v {
			return true
		}
	}
	return false
}
