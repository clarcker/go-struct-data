package 栈

import "fmt"

type ListStack struct {
	top  int
	list []int
	max  int
}

func (l *ListStack) InitStack(max int) {
	if max >= 0 {
		l.max = max
		l.top = 0
		l.list = make([]int, max)
	}
}

func (l ListStack) StackEmpty() bool {
	return l.top == 0
}

func (l *ListStack) Push(x int) bool {
	if l.top >= l.max {
		return false
	}
	l.list[l.top] = x
	l.top++
	return true
}

func (l *ListStack) Pop(x *int) bool {
	if l.top == 0 {
		return false
	}
	l.top--
	*x = l.list[l.top]
	return true
}

func (l ListStack) GetTop(x *int) bool {
	if l.top == 0 {
		return false
	}
	*x = l.list[l.top-1]
	return true
}

func (l *ListStack) Destroy() {
	l.list = nil
	l.max = 0
	l.top = 0
}

func (l ListStack) Print(s, fmts string) {
	if s != "" {
		fmt.Println(s)
	}
	for i := l.top - 1; i >= 0; i-- {
		fmt.Printf("%v%v", l.list[i], fmts)
	}
	fmt.Println()
}
