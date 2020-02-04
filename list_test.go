package golink

import (
	"testing"
	"fmt"
)

func TestArrayList(t *testing.T) {
	l := NewArrayList()
	l.Add(1)
	l.Add(2)
	l.Add(3)
	fmt.Println("******************")
	l.OrderRead()
	l.PostRead()
	// 在指定位置插入元素
	fmt.Println("******************")
	l.Insert(9,2)
	l.OrderRead()
	l.PostRead()
	l.Insert(8,1)
	fmt.Println("******************")
	l.OrderRead()
	l.PostRead()
	l.Insert(6,5)
	fmt.Println("******************")
	l.OrderRead()
	l.PostRead()
	l.Insert(11,7)
	fmt.Println("******************")
	l.OrderRead()
	l.PostRead()
}
