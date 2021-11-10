package 线性表

import "fmt"

type circularNode struct {
	next *circularNode
	val  int
}
type CircularLinear struct {
	head    *circularNode
	len     int
	maxSize int
}

func (c *CircularLinear) Init(lens int) {
	c.maxSize = lens
}

func (c *CircularLinear) InsertEnd(e int) bool {
	if c.maxSize < c.len+1 {
		return false
	}
	p := new(circularNode)
	p.val, p.next = e, c.head
	if c.len == 0 {
		c.head = p
		p.next = p
	} else {
		q := c.head
		for q.next != c.head {
			q = q.next
		}
		q.next = p
	}
	c.len++
	return true
}

func (c *CircularLinear) InsertLocate(index int, e int) bool {
	if index < 0 || index > c.len || c.len+1 > c.maxSize {
		return false
	}
	return true
}

func (c CircularLinear) LocateElem(e int) int {
	panic("implement me")
}

func (c CircularLinear) GetElem(index int) (int, bool) {
	panic("implement me")
}

func (c CircularLinear) DeleteElem(e int) bool {
	panic("implement me")
}

func (c CircularLinear) DeleteIndex(index int) bool {
	panic("implement me")
}

func (c CircularLinear) Empty() bool {
	panic("implement me")
}

func (c CircularLinear) Length() int {
	panic("implement me")
}

func (c CircularLinear) Print(s string) {
	for p, flag := c.head, true; p != nil && ((p == c.head && flag) || (p != c.head && !flag)); p = p.next {
		flag = false
		fmt.Printf("%v%v", p.val, s)
	}
	fmt.Println()
}
func (c CircularLinear) Destroy() {
	panic("implement me")
}

func (c *CircularLinear) DeleteAll() {
	for c.len > 0 {
		if c.len == 1 {
			fmt.Println(c.head.val)
			c.head = nil
			c.len--
		}
		var min = new(circularNode)
		var pre = new(circularNode)
		min.next, min.val = c.head, c.head.val
		pre = c.head
		for p, q := c.head, c.head.next; q != c.head; q = q.next {
			if q.val < min.val {
				pre = p
				min = q
			}
		}
		if min == pre {

		}
	}
}
