package 栈

import "fmt"

type nodeList struct {
	val  int
	next *nodeList
}

type NodeStack struct {
	max, count int
	head       *nodeList
}

func (n *NodeStack) InitStack(max int) {
	n.head, n.count, n.max = nil, 0, max
}

func (n NodeStack) StackEmpty() bool {
	return n.count == 0
}

func (n *NodeStack) Push(x int) bool {
	if n.count >= n.max && n.max != 0 {
		return false
	}
	n.count++
	p := new(nodeList)
	p.val, p.next = x, n.head
	n.head = p
	return true
}

func (n *NodeStack) Pop(x *int) bool {
	if n.count == 0 {
		return false
	}
	n.count--
	*x = n.head.val
	n.head = n.head.next
	return true
}

func (n NodeStack) GetTop(x *int) bool {
	if n.count == 0 {
		return false
	}
	*x = n.head.val
	return true
}

func (n NodeStack) Destroy() {
	n.head, n.max, n.count = nil, 0, 0
}

func (n NodeStack) Print(s, fmts string) {
	fmt.Println(s)
	for p := n.head; p != nil; p = p.next {
		fmt.Printf("%v%v", p.val, fmts)
	}
	fmt.Println()
}
