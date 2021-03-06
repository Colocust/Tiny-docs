package strcture

import "math"

type (
	ListNode struct {
		value interface{}
		prev  *ListNode
		next  *ListNode
	}

	List struct {
		head *ListNode
		tail *ListNode
		len  uint
	}
)

func NewList() *List {
	return &List{
		len: 0,
	}
}

func newNode(value interface{}) *ListNode {
	return &ListNode{
		value: value,
	}
}

func (l *List) AddHead(value interface{}) {
	node := newNode(value)
	if l.len == 0 {
		l.head, l.tail = node, node
	} else {
		node.next, l.head.prev, l.head = l.head, node, node
	}
	l.len++
}

func (l *List) AddTail(value interface{}) {
	node := newNode(value)
	if l.len == 0 {
		l.head, l.tail = node, node
	} else {
		node.prev, l.tail.next, l.tail = l.tail, node, node
	}
	l.len++
}

//将尾节点移到头节点
func (l *List) RotateTailToHead() {
	if l.len <= 1 {
		return
	}
	//取出尾节点
	tail := l.tail
	//调整尾节点前驱节点指针信息以及链表现尾节点
	l.tail = tail.prev
	l.tail.next = nil
	//将尾节点移动到头节点
	l.head.prev, tail.next, tail.prev = tail, l.head, nil
	l.head = tail
}

//将头节点移到尾节点
func (l *List) RotateHeadToTail() {
	if l.len <= 1 {
		return
	}

	head := l.head
	l.head = head.next
	l.head.prev = nil

	l.tail.next, head.prev, head.next = head, l.tail, nil
	l.tail = head
}

//在链表中查找指定value的节点
func (l *List) Search(value interface{}) *ListNode {
	cur := l.head

	for cur != nil {
		if cur.value == value {
			return cur
		}
		cur = cur.next
	}

	return nil
}

//获取指定索引的节点
func (l *List) Index(index int) *ListNode {
	if index < 0 {
		abs := uint(math.Abs(float64(index)))
		if abs > l.Len() {
			return nil
		}
	} else {
		if uint(index) > l.len-1 {
			return nil
		}
	}

	var n *ListNode
	if index < 0 {
		index = (-index) - 1
		n = l.tail
		for index > 0 && n != nil {
			n = n.prev
			index--
		}
	} else {
		n = l.head
		for index > 0 && n != nil {
			n = n.next
			index--
		}
	}
	return n
}

//指定节点前后插入一个新节点 after决定前后
func (l *List) Insert(oldNode *ListNode, value interface{}, after bool) {
	node := newNode(value)

	if after {
		//先更新新节点前后指针
		node.prev, node.next = oldNode, oldNode.next
		if l.tail == oldNode {
			l.tail = node
		}
	} else {
		node.next, node.prev = oldNode, oldNode.prev
		if l.head == oldNode {
			l.head = node
		}
	}

	//更新oldNode以及oldNode.next/prev
	if node.prev != nil {
		node.prev.next = node
	}

	if node.next != nil {
		node.next.prev = node
	}

	l.len++
}

//如果删除的节点不在指定的list 需要开发者自行注意
func (l *List) Del(n *ListNode) {
	if n.prev == nil {
		l.head = n.next
	} else {
		n.prev.next = n.next
	}

	if n.next == nil {
		l.tail = n.prev
	} else {
		n.next.prev = n.prev
	}

	l.len--
}

func (l *List) Len() uint {
	return l.len
}

func (l *List) Head() *ListNode {
	return l.head
}

func (l *List) Tail() *ListNode {
	return l.tail
}

func (n *ListNode) Value() interface{} {
	return n.value
}

func (n *ListNode) Prev() *ListNode {
	return n.prev
}

func (n *ListNode) Next() *ListNode {
	return n.next
}
