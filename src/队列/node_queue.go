package 队列

import "fmt"

type nodeList struct {
	val  int
	next *nodeList
}

type NodeQueue struct {
	tail *nodeList
	head *nodeList
	len  int
}

func (n NodeQueue) InitQueue(max int) {
	panic("implement me")
}

func (n NodeQueue) QueueEmpty() bool {
	return n.len == 0
}

func (n *NodeQueue) EnQueue(x int) bool {
	q := new(nodeList)
	q.next, q.val = nil, x
	if n.len == 0 {
		n.head = q
	} else {
		n.tail.next = q
	}
	n.tail = q
	n.len++
	return true
}

func (n *NodeQueue) DeQueue(x *int) bool {
	if n.len == 0 {
		return false
	}
	*x = n.head.val
	n.head = n.head.next
	n.len--
	return true
}

func (n NodeQueue) GetTop(x *int) bool {
	if n.len == 0 {
		return false
	}
	*x = n.head.val
	n.head = n.head.next
	return true
}

func (n NodeQueue) Destroy() {
	panic("implement me")
}

func (n NodeQueue) Print(s, fmts string) {
	fmt.Println(s)
	for q := n.head; q != nil; q = q.next {
		fmt.Printf("%v%v", q.val, fmts)
	}
	fmt.Println()
}
