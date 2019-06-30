package goutil

import (
	"fmt"
	"reflect"
	"sort"
)

//Array 仿php中的索引数组 相当于带顺序可排序的map
//通过下标取值方法   Array.Body[Array.Keys[n]]
//通过key取值的方法  Array.Body[n]
type Array struct {
	Body         map[string]interface{} //原map
	Keys         []string               //slice有序化的key
	ParamForSort map[string]interface{} //排序用到的外部变量
	sortFunc     ArraySortFuncType      //排序方法
}

//NewArray 传入一个map 默认key asc排序
func NewArray(data ...map[string]interface{}) *Array {
	if len(data) == 0 {
		return new(Array)
	}
	var Keys []string
	for k := range data[0] {
		Keys = append(Keys, k)
	}
	sort.Strings(Keys)

	res := Array{
		Body: data[0],
		Keys: Keys,
		//默认给一个 desc排序方法
		sortFunc: func(arr *Array, i, j int) bool {
			return arr.Keys[i] > arr.Keys[j]
		},
	}
	return &res
}

//Insert 插入
//k 索引名 v 数据  i 插入的位置
func (p *Array) Insert(k string, v interface{}, i int) error {
	if len(p.Keys) < i {
		return fmt.Errorf("超出了Array的范围")
	}
	p.Body[k] = v
	if len(p.Keys) > i {
		tmp := append(p.Keys[:i], k)
		p.Keys = append(tmp, p.Keys[i:]...)
		return nil
	} else {
		p.Keys = append(p.Keys, k)
	}
	return nil
}

//Delete 删除
//v 索引名 string 或 位置 int
func (p *Array) Delete(v interface{}) (err error) {
	var i int
	var k string
	var ok bool
	if k, ok = v.(string); ok {
		if i, err = p.GetIndexFromKey(k); err != nil {
			return
		}
		delete(p.Body, k)
		if len(p.Keys) == i+1 {
			p.Keys = p.Keys[:i]
		} else {
			p.Keys = append(p.Keys[:i], p.Keys[i+1:]...)
		}
	} else if i, ok = v.(int); ok {
		if len(p.Keys) > i {
			delete(p.Body, p.Keys[i])
			if len(p.Keys) == i+1 {
				p.Keys = p.Keys[:i]
			} else {
				p.Keys = append(p.Keys[:i], p.Keys[i+1:]...)
			}
		}
		err = fmt.Errorf("超出了Array的范围")
	} else {
		err = fmt.Errorf("只接受string 或 int")
	}
	return
}

//GetIndexFromKey 从索引获取下标
func (p *Array) GetIndexFromKey(k string) (i int, err error) {
	var v string
	for i, v = range p.Keys {
		if v == k {
			return
		}
	}
	err = fmt.Errorf("索引不存在: %s", k)
	return
}

//InArray 判断是否在数组中
func (p *Array) InArray(value interface{}) bool {
	if len(p.Keys) == 0 {
		return false
	}
	if reflect.TypeOf(value) != reflect.TypeOf(p.Body[p.Keys[0]]) {
		return false
	}
	for _, v := range p.Body {
		if value == v {
			return true
		}
	}
	return false
}

//Foreach 指定遍历方法
//回调函数接受 索引 下标 数据
func (p *Array) Foreach(f func(k string, i int, v interface{}) error) error {
	for index, key := range p.Keys {
		if err := f(key, index, p.Body[key]); err != nil {
			return err
		}
	}
	return nil
}

//ArraySortFuncType 自定义排序方法类型
/**
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

//Sort 自定义排序方法
func (p *Array) Sort(sortFunc ArraySortFuncType) {
	p.sortFunc = sortFunc
	sort.Sort(p)
}
