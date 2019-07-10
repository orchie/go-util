package goutil

import "fmt"

//Tree 树形结构
type Tree struct {
	ID       int
	Pid      int
	Children []TreeItf
}

//TreeItf 树形结构接口
type TreeItf interface {
	GetID() int
	GetPID() int
	GetChildren() []TreeItf
	SetChildren([]TreeItf)
}

//GetID GetID
func (p *Tree) GetID() int {
	return p.ID
}

//GetPID GetPID
func (p *Tree) GetPID() int {
	return p.Pid
}

//GetChildren GetChildren
func (p *Tree) GetChildren() []TreeItf {
	return p.Children
}

//SetChildren SetChildren
func (p *Tree) SetChildren(t []TreeItf) {
	p.Children = t
}

//SliceToTree 将slice转化为Tree
func SliceToTree(s []TreeItf) []TreeItf {
	return sliceToTreeLoop(&s, 0)
}

//SliceToTree 递归函数
var depth int

//使用指针是为了保证递归操作的s是同一个(非指针传值就拷贝为另一个指针了)
func sliceToTreeLoop(s *[]TreeItf, pid int) (t []TreeItf) {
	// fmt.Printf("depth:%d s 递归长度 %d,pid:%d\n", depth, len(*s), pid)
	//删除元素后 下一个元素会移到当前,下次遍历就会漏掉 所以用另个变量转存下标
	realK := 0
	for _ = range *s {
		// fmt.Printf("k:%d,id:%d,pid:%d,nowPid:%d ;", realK, (*s)[realK].GetID(), (*s)[realK].GetPID(), pid)
		if (*s)[realK].GetPID() == pid {
			t = append(t, (*s)[realK])
			//删前留名
			vv := (*s)[realK]
			//这里删除元素
			if realK == len(*s)-1 {
				*s = (*s)[:realK]
			} else if realK < len(*s)-1 {
				*s = append((*s)[:realK], (*s)[realK+1:]...)
				realK--
			} else {
				continue
			}
			realK++
			depth++
			//使用v 在原 TreeItf 上进行操作
			vv.SetChildren(sliceToTreeLoop(s, vv.GetID()))
			depth--
			// fmt.Printf("id:%d over s 长度 %d\n", vv.GetID(), len(*s))
		}
	}
	fmt.Printf("\n")
	return t
}
