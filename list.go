package golink

// list列表结构体
type Node struct {
	item  interface{}
	next  *Node
	pre   *Node
}
type arrayList struct {
	head *Node
	lengh int
}
// NewArrayList
func NewArrayList() *arrayList {
	return &arrayList{
		head: nil,
		lengh:0,
	}
}

// 向arraylist中添加元素
func (l *arrayList) Add(item interface{}) {
	if l == nil {
		panic("list is nil")
	}
	if l.lengh == 0 {
		node :=&Node{item,nil,nil}
		node.pre = node
		node.next = node
		l.lengh++
		l.head = node
		return
	}
	p := l.head
	for p.next != l.head {
		p = p.next
	}
	p.next = &Node{
		item:  item,
		next:  l.head,
		pre:   p,
	}
	l.head.pre = p.next
	l.lengh++
}

// 在指定位置插入元素
func (l *arrayList)Insert(item interface{},index int){
	if index > l.lengh{
		l.Add(item)
		return
	}
	pre := l.head
	next := l.head.next
	for i:=0;i<index-1;i++{
		pre = pre.next
		next = next.next
	}
	head :=&Node{
		item:pre.item,
		next:next,
		pre:pre,
	}
	pre.next = head
	pre.item = item
	next.pre = head
	l.lengh++
}
// 顺序遍历
func (l *arrayList) OrderRead()(result []interface{} ){
	p := l.head
	for p.next != l.head {
		result = append(result,p.item)
		p = p.next
	}
	result = append(result,p.item)
	return
}

// 倒叙遍历
func (l *arrayList) PostRead() (result []interface{} ){
	p :=l.head.pre
	for p !=l.head{
		result = append(result,p.item)
		p = p.pre
	}
	result = append(result,p.item)
	return
}
