package golink

import "fmt"

// list列表结构体
type list struct {
	item  interface{}
	next  *list
	pre   *list
	lengh int
}

// NewArrayList
func NewArrayList() *list {
	return &list{
		lengh: 0,
	}
}

// 向arraylist中添加元素
func (l *list) Add(item interface{}) {
	if l == nil {
		panic("list is nil")
	}
	if l.lengh == 0 {
		l.item = item
		l.pre = l
		l.next = l
		l.lengh++
		return
	}
	p := l
	for p.next != l {
		p = p.next
	}
	p.next = &list{
		item:  item,
		next:  l,
		pre:   p,
		lengh: l.lengh + 1,
	}
	l.pre = p.next
}

// 在指定位置插入元素
func (l *list)Insert(item interface{},index int){
	pre := l
	next := l.next
	for i:=0;i<index-1;i++{
		pre = pre.next
		next = next.next
	}
	head :=&list{
		item:pre.item,
		next:next,
		pre:pre,
		lengh:l.lengh+1,
	}
	pre.next = head
	pre.item = item
	next.pre = head
}
// 顺序遍历
func (l *list) OrderRead() {
	p := l
	for p.next != l {
		fmt.Print(p.item," ")
		p = p.next
	}
	fmt.Println(p.item)
}

// 倒叙遍历
func (l *list) PostRead() {
	p :=l.pre
	for p !=l{
		fmt.Print(p.item," ")
		p = p.pre
	}
	fmt.Println(p.item)
}
