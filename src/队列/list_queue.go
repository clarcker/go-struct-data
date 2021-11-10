package 队列

import (
	"fmt"
)

type CircularQueue struct {
	head, tail, max, len int
	list                 []int
}

func (c *CircularQueue) InitQueue(max int) {
	c.tail, c.head, c.max, c.len, c.list = 0, 0, max, 0, make([]int, max)
}

func (c CircularQueue) QueueEmpty() bool {
	return c.len == 0
}

func (c *CircularQueue) EnQueue(x int) bool {
	if c.len == c.max {
		return false
	}
	c.list[c.tail] = x
	c.tail = (c.tail + 1) % c.max
	c.len++
	return true
}

func (c *CircularQueue) DeQueue(x *int) bool {
	if c.len == 0 {
		return false
	}
	*x = c.list[c.head]
	c.head = (c.head + 1) % c.max
	c.len--
	return true
}

func (c CircularQueue) GetTop(x *int) bool {
	if c.len == 0 {
		return false
	}
	*x = c.list[c.head]
	return true
}

func (c *CircularQueue) Destroy() {
	c.tail, c.head, c.max, c.len, c.list = 0, 0, 0, 0, nil
}

func (c CircularQueue) Print(s, fmts string) {
	fmt.Println(s)
	for i := 0; i < c.len; i++ {
		fmt.Printf("%v%v", c.list[(c.head+i)%c.max], fmts)
	}
	fmt.Println()
}

type CircularQueue2 struct {
	head, tail, max, tag int
	list                 []int
}

func (c *CircularQueue2) InitQueue(max int) {
	c.tag, c.head, c.tail, c.max, c.list = 0, 0, 0, max, make([]int, max)
}

func (c CircularQueue2) QueueEmpty() bool {
	return c.tag == 0 && c.tail == c.head
}

func (c *CircularQueue2) EnQueue(x int) bool {
	if c.tag == 1 && c.tail == c.head {
		return false
	}
	c.list[c.tail] = x
	c.tail = (c.tail + 1) % c.max
	if c.tail == c.head {
		c.tag = 1
	}
	return true
}

func (c *CircularQueue2) DeQueue(x *int) bool {
	if c.tail == c.head && c.tag == 0 {
		return false
	}
	*x = c.list[c.head]
	c.head = (c.head + 1) % c.max
	if c.head == c.tail {
		c.tag = 0
	}
	return true
}

func (c CircularQueue2) GetTop(x *int) bool {
	panic("implement me")
}

func (c CircularQueue2) Destroy() {
	panic("implement me")
}

func (c CircularQueue2) Print(s, fmts string) {
	fmt.Println(s)
	if c.tail == c.head && c.tag == 0 {
		return
	}
	if c.tail == c.head && c.tag == 1 {
		for i := 0; i < c.max; i++ {
			fmt.Printf("%v%v", c.list[(i+c.head)%c.max], fmts)
		}
	} else {
		for i := c.head; i%c.max != c.tail; i++ {
			fmt.Printf("%v%v", c.list[i%c.max], fmts)
		}
	}

	fmt.Println()
}
