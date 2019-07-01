package goutil

import (
	"reflect"
	"sort"
)

//排序容器
//对slice结构体排序
type bodyWrapper struct {
	Body     []interface{}
	SortFunc SortFuncType
}

//SortFuncType 排序方法
type SortFuncType func(p, q *interface{}) bool

//实现sort接口
func (bw bodyWrapper) Len() int {
	return len(bw.Body)
}

func (bw bodyWrapper) Swap(i, j int) {
	bw.Body[i], bw.Body[j] = bw.Body[j], bw.Body[i]
}

func (bw bodyWrapper) Less(i, j int) bool {
	return bw.SortFunc(&bw.Body[i], &bw.Body[j])
}

//Sort 外部使用 通用排序方法
func Sort(body []interface{}, sortFunc SortFuncType) {
	sort.Sort(bodyWrapper{body, sortFunc})
}

//SortField 指定字段排序
func SortField(body []interface{}, field string) {
	Sort(body, func(p, q *interface{}) bool {
		v := reflect.ValueOf(*p)
		i := v.FieldByName(field)
		v = reflect.ValueOf(*q)
		j := v.FieldByName(field)
		return i.String() < j.String()
	})
}
