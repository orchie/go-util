package goutil

import (
	"sort"
)

//仿php中的数组
//Body保存 原map
//Keys 保存 slice有序化的key
//通过下标取值方法   Array.Body[Array.Keys[n]]
//通过key取值的方法  Array.Body[n]
type Array struct {
	Body     map[string]interface{}
	Keys     []string
	sortFunc ArraySortFuncType
}

//传入一个map 默认key asc排序
func NewArray(data map[string]interface{}) *Array {
	var Keys []string
	for k, _ := range data {
		Keys = append(Keys, k)
	}
	sort.Strings(Keys)

	res := Array{
		Body: data,
		Keys: Keys,
		//默认给一个 desc排序方法
		sortFunc: func(arr *Array, i, j int) bool {
			return arr.Keys[i] > arr.Keys[j]
		},
	}
	return &res
}

//Array 无法解决 for...range  问题 无法用for...range
//指定遍历方法
func (p *Array) Foreach(f func(k string, i int, v interface{}) error) error {
	for index, key := range p.Keys {
		if err := f(key, index, p.Body[key]); err != nil {
			return err
		}
	}
	return nil
}

/**
	自定义排序方法类型
	arr 当前Array数组
	i,j 排序的索引
	extends 额外用到的参数 不用就传空
**/
type ArraySortFuncType func(arr *Array, i, j int) bool

//实现排序接口
func (p *Array) Len() int {
	return len(p.Keys)
}

func (p *Array) Swap(i, j int) {
	p.Keys[i], p.Keys[j] = p.Keys[j], p.Keys[i]
}

func (p *Array) Less(i, j int) bool {
	return p.sortFunc(p, i, j)
}

//外部使用 自定义排序方法
func (p *Array) Sort(sortFunc ArraySortFuncType) {
	p.sortFunc = sortFunc
	sort.Sort(p)
}

//使用 slice 来实现数组
//先不用 留着这个思路
type Arr []Kv

type Kv struct {
	K string
	V interface{}
}

//传入一个map 默认key asc排序
func NewArr(data map[string]interface{}) Arr {
	var res Arr
	var KeysArr []string
	for k, _ := range data {
		KeysArr = append(KeysArr, k)
	}
	sort.Strings(KeysArr)
	for _, v := range KeysArr {
		res = append(res, Kv{
			K: v,
			V: data[v],
		})
	}
	return res
}

//获取指定index
func (p Arr) get(key string) interface{} {
	for _, v := range p {
		if v.K == key {
			return v.V
		}
	}
	return false
}
