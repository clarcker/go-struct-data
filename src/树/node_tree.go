package 树

import "fmt"

const MAXTREENODE = 100

type NodeTree struct {
	leftChild  *NodeTree
	rightChild *NodeTree
	val        int
}

type nodeList struct {
	val  *NodeTree
	next *nodeList
}

type NodeQueue struct {
	tail *nodeList
	head *nodeList
	len  int
}

func (n *NodeQueue) EnQueue(x *NodeTree) bool {
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

func (n *NodeQueue) DeQueue(x **NodeTree) bool {
	if n.len == 0 {
		return false
	}
	*x = n.head.val
	n.head = n.head.next
	n.len--
	return true
}

func (n NodeQueue) GetTop(x **NodeTree) bool {
	if n.len == 0 {
		return false
	}
	*x = n.head.val
	n.head = n.head.next
	return true
}

func (n NodeQueue) QueueEmpty() bool {
	return n.len == 0
}

func (n *NodeTree) ListTreeToNodeTree(arr []int) {
	var queue NodeQueue
	flag := 0
	n.val = arr[1]
	p := n
	queue.EnQueue(p)
	for i := 2; i < arr[0]; i++ {
		switch i % 2 {
		case 0:
			{
				if arr[i] != 0 {
					flag = 1
					queue.GetTop(&p)
					p.leftChild = new(NodeTree)
					p.leftChild.val = arr[i]
					queue.EnQueue(p.leftChild)
				} else {
					flag = 0
				}
			}
		case 1:
			{
				if arr[i] != 0 {
					queue.DeQueue(&p)
					p.rightChild = new(NodeTree)
					p.rightChild.val = arr[i]
					queue.EnQueue(p.rightChild)
				} else {
					if flag == 1 {
						queue.DeQueue(&p)
					}
				}
			}
		}
	}
}

func (n *NodeTree) PreVisitRecursion() {
	for n != nil {
		fmt.Printf("%v ", n.val)
		n.leftChild.PreVisitRecursion()
		n.rightChild.PreVisitRecursion()
		return
	}
}

func (n *NodeTree) MidVisitRecursion() {
	for n != nil {
		n.leftChild.MidVisitRecursion()
		fmt.Printf("%v ", n.val)
		n.rightChild.MidVisitRecursion()
		return
	}
}

func (n *NodeTree) PostVisitRecursion() {
	for n != nil {
		n.leftChild.PostVisitRecursion()
		n.rightChild.PostVisitRecursion()
		fmt.Printf("%v ", n.val)
		return
	}
}

func (n *NodeTree) AddTreeNode(isLeft bool, val int) {
	newP := new(NodeTree)
	newP.val = val
	if isLeft {
		n.leftChild = newP
	} else {
		n.rightChild = newP
	}
}

//
type nodeStack struct {
	val  *NodeTree
	next *nodeStack
}

type NodeStack struct {
	max, count int
	head       *nodeStack
}

func (n *NodeStack) InitStack(max int) {
	n.head, n.count, n.max = nil, 0, max
}

func (n NodeStack) StackEmpty() bool {
	return n.count == 0
}

func (n *NodeStack) Push(x *NodeTree) bool {
	if n.count >= n.max && n.max != 0 {
		return false
	}
	n.count++
	p := new(nodeStack)
	p.val, p.next = x, n.head
	n.head = p
	return true
}

func (n *NodeStack) Pop(x **NodeTree) bool {
	if n.count == 0 {
		return false
	}
	n.count--
	*x = n.head.val
	n.head = n.head.next
	return true
}

func (n NodeStack) GetTop(x **NodeTree) bool {
	if n.count == 0 {
		return false
	}
	*x = n.head.val
	return true
}

func (n *NodeTree) PreVisit() {
	p := n
	var stack NodeStack
	stack.InitStack(MAXTREENODE)
	stack.Push(p)
	for !stack.StackEmpty() || p != nil {
		if p != nil {
			fmt.Printf("%v ", p.val)
			if p != n {
				stack.Push(p)
			}
			p = p.leftChild
		} else {
			stack.Pop(&p)
			p = p.rightChild
		}
	}
	fmt.Println()
}

func (n *NodeTree) MidVisit() {
	p := n
	var stack NodeStack
	stack.InitStack(MAXTREENODE)
	stack.Push(p)
	for !stack.StackEmpty() || p != nil {
		if p != nil {
			if p != n {
				stack.Push(p)
			}
			p = p.leftChild
		} else {
			stack.Pop(&p)
			fmt.Printf("%v ", p.val)
			p = p.rightChild
		}
	}
	fmt.Println()
}

func (n NodeTree) PostVisit() {

}

func (n *NodeTree) LevelVisit() {
	var queue NodeQueue
	queue.EnQueue(n)
	q := n
	for !queue.QueueEmpty() {
		queue.DeQueue(&q)
		fmt.Printf("%v ", q.val)
		if q.leftChild != nil {
			queue.EnQueue(q.leftChild)
		}
		if q.rightChild != nil {
			queue.EnQueue(q.rightChild)
		}
	}
	fmt.Println()
}
