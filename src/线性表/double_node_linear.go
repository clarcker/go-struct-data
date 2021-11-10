package 线性表

import "fmt"

type doubleNode struct {
	previous *doubleNode
	val      int
	next     *doubleNode
}

type DoubleNodeLinear struct {
	head *doubleNode
	len  int
	end  *doubleNode
}

func (d DoubleNodeLinear) Init(lens int) {
	panic("implement me")
}

func (d *DoubleNodeLinear) InsertHead(e int) bool {
	p := new(doubleNode)
	p.val = e
	if d.len == 0 {
		p.next, p.previous = nil, nil
		d.head, d.end = p, p
	} else {
		p.next, p.previous = d.head, nil
		d.head.previous = p
		d.head = p
	}
	d.len++
	return true
}

func (d *DoubleNodeLinear) InsertEnd(e int) bool {
	p := new(doubleNode)
	p.val = e
	if d.len == 0 {
		p.next, p.previous = nil, nil
		d.head, d.end = p, p
	} else {
		p.previous, p.next = d.end, nil
		d.end.next, d.end = p, p
	}
	d.len++
	return true
}

func (d *DoubleNodeLinear) InsertLocate(index int, e int) bool {
	if index < 0 || index > d.len {
		return false
	}
	if index == d.len {
		q := new(doubleNode)
		q.next = nil
		q.previous = d.end
		q.val = e
		d.end.next = q
		d.end = q
		d.len++
	}
	for count, p := 0, d.head; count <= index; count++ {
		if count == index {
			q := new(doubleNode)
			q.next = p
			q.val = e
			q.previous = p.previous
			if index == 0 {
				d.head = q
			} else {
				p.previous.next = q
			}
			p.previous = q
			d.len++
			return true
		}
	}
	return false
}

func (d DoubleNodeLinear) LocateElem(e int) int {
	for count, p := 0, d.head; p != nil; count++ {
		if p.val == e {
			return count
		}
		p = p.next
	}
	return -1
}

func (d DoubleNodeLinear) GetElem(index int) (int, bool) {
	if index < 0 || index >= d.len {
		return -1, false
	}
	for count, p := 0, d.head; count <= index; count++ {
		if count == index {
			return p.val, true
		}
		p = p.next
	}
	return -1, false
}

func (d *DoubleNodeLinear) DeleteElem(e int) bool {
	for p := d.head; p != nil; p = p.next {
		if p.val == e {
			if d.head == p {
				d.head = p.next
				p.next.previous = nil
			} else if d.end == p {
				p.previous.next = nil
				d.end = p.previous
			} else {
				p.previous.next = p.next
				p.next.previous = p.previous
			}
			d.len--
			return true
		}
	}
	return false
}

func (d *DoubleNodeLinear) DeleteIndex(index int) bool {
	if index < 0 || index >= d.len {
		return false
	}
	for count, p := 0, d.head; p != nil; count++ {
		if count == index {
			if index == 0 {
				d.head = p.next
				p.next.previous = nil
			} else if index == d.len-1 {
				p.previous.next = nil
				d.end = p.previous
			} else {
				p.previous.next = p.next
				p.next.previous = p.previous
			}
			d.len--
			return true
		}
		p = p.next
	}
	return false
}

func (d DoubleNodeLinear) Empty() bool {
	return d.len == 0
}

func (d DoubleNodeLinear) Length() int {
	return d.len
}

func (d DoubleNodeLinear) Print(s string) {
	for p := d.head; p != nil; p = p.next {
		fmt.Printf("%v%s", p.val, s)
	}
	fmt.Println()
	for q := d.end; q != nil; q = q.previous {
		fmt.Printf("%v%s", q.val, s)
	}
	fmt.Println()
}
func (d *DoubleNodeLinear) Destroy() {
	d.head, d.end, d.len = nil, nil, 0
}
