package 线性表

import (
	"fmt"
)

type node struct {
	val  int
	next *node
}

type NodeList struct {
	head *node
	len  int
}

func (n NodeList) Init(lens int) {
}

func (n *NodeList) InsertEnd(e int) bool {
	n.len++
	if n.head == nil {
		n.head = new(node)
		n.head.val = e
		return true
	}
	p := n.head
	for p.next != nil {
		p = p.next
	}
	p.next = new(node)
	p.next.val = e
	return true
}

func (n *NodeList) InsertHead(e int) bool {
	p := new(node)
	p.next = n.head
	p.val = e
	n.head = p
	n.len++
	return true
}

func (n *NodeList) InsertLocate(index int, e int) bool {
	if index < 0 || index > n.len {
		return false
	}
	n.len++
	p := n.head
	if index == 0 {
		p = new(node)
		p.val = e
		p.next = n.head
		n.head.next = nil
		n.head = p
		return true
	}
	i := 0
	for ; i <= index && p.next != nil; i++ {
		if i == index {
			q := new(node)
			q.val = p.val
			q.next = p.next
			p.val = e
			p.next = q
		}
		p = p.next
	}
	panic("implement me")
}

func (n NodeList) LocateElem(e int) int {
	p := n.head
	for count := 0; p != nil; count++ {
		if p.val == e {
			return count
		}
		p = p.next
	}
	return -1
}

func (n NodeList) GetElem(index int) (int, bool) {
	if index < 0 || index >= n.len {
		return -1, false
	}
	for p, count := n.head, 0; p != nil; count++ {
		if count == index {
			return p.val, true
		}
		p = p.next
	}
	return -1, false
}

func (n *NodeList) DeleteElem(e int) bool {
	if n.len < 1 {
		return false
	}
	if n.head.val == e {
		n.head = n.head.next
		n.len--
		return true
	}
	var beforeP *node
	for p := n.head; p != nil; p = p.next {
		if p.val == e {
			beforeP.next = p.next
			n.len--
			return true
		}
		beforeP = p
	}
	return false
}

func (n *NodeList) DeleteIndex(index int) bool {
	if index < 0 || index >= n.len {
		return false
	}
	if index == 0 {
		n.head = n.head.next
		n.len--
		return true
	}
	for count, p := 0, n.head; p != nil; count++ {
		if count == index-1 {
			p.next = p.next.next
			n.len--
			return true
		}
		p = p.next
	}
	return false
}

func (n NodeList) Empty() bool {
	return n.len == 0
}

func (n NodeList) Length() int {
	return n.len
}

func (n NodeList) Print(s string) {
	for p := n.head; p != nil; p = p.next {
		fmt.Printf("%v%s", p.val, s)
	}
	fmt.Println()
}
func (n *NodeList) Destroy() {
	n.head = nil
	n.len = 0
}

func (n *NodeList) DeleteAllX(x int) {
	if n.head == nil {
		return
	}
	if n.head.val == x {
		n.head = n.head.next
		n.len--
	}
	if n.len < 2 {
		return
	}
	for p, q := n.head, n.head.next; q != nil; q = q.next {
		if q.val == x {
			p.next = q.next
			n.len--
		} else {
			p = q
		}
	}
}

func (n *NodeList) InsertSort(x int) {
	if n.len == 0 {
		n.head = new(node)
		n.head.val = x
	} else {
		for p := n.head; p != nil; p = p.next {
			if p.val > x {
				q := new(node)
				q.val, q.next = p.val, p.next
				p.next, p.val = q, x
				break
			}
			if p.next == nil {
				q := new(node)
				q.val = x
				p.next = q
				break
			}
		}
	}
	n.len++
}

func (n *NodeList) DeleteAllSame() {
	if n.len < 2 {
		return
	}
	for p, q := n.head, n.head.next; q != nil; q = q.next {
		if p.val == q.val {
			p.next = q.next
			n.len--
		} else {
			p = q
		}
	}
}

func (n *NodeList) P39T25() {
	count := 0
	var q *node
	p := n.head
	for ; count < (n.len-1)/2; count++ {
		p = p.next
	}
	q = p.next
	p.next = nil
	n.Print(" ")
	var temp = new(NodeList)
	temp.head = q
	q = q.next
	temp.head.next = nil
	for q != nil {
		tem2 := q.next
		q.next = temp.head
		temp.head = q
		q = tem2
	}
	temp.Print(" ")
	for p, q = n.head, temp.head; q != nil && p != nil; {
		temp1, temp2 := p.next, q.next
		p.next = q
		q.next = temp1
		p, q = temp1, temp2
	}
}
