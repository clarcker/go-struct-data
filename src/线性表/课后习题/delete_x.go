package 课后习题

import "fmt"

/*
 */
type NodeList struct {
	val  int
	next *NodeList
}

func InsertHead(l **NodeList, e int) {
	if l == nil {
		*l = new(NodeList)
		(*l).val = e
	}
	p := new(NodeList)
	p.val, p.next = e, *l
	*l = p
}
func PrintNode(l *NodeList) {
	for p := l; p != nil; p = p.next {
		fmt.Printf("%v ", p.val)
	}
	fmt.Println()
}

func DeleteX(l *NodeList, x int) {
	if l == nil {
		return
	}
	if l.val == x {
		l = l.next
		DeleteX(l, x)
	} else {
		DeleteX(l, x)
	}
}
