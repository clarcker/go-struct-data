package 栈

import "fmt"

type ShareStack struct {
	max, top1, top2 int
	list            []int
}

func (s *ShareStack) InitStack(max int) {
	s.max, s.top1, s.top2, s.list = max, 0, max-1, make([]int, max)
}

func (s ShareStack) StackEmpty(stackName int) bool {
	if stackName == 0 {
		return s.top1 == 0
	} else {
		return s.top2 == s.max-1
	}
}

func (s *ShareStack) Push(stackName, x int) bool {
	if s.top1 == s.top2+1 {
		return false
	}
	if stackName == 0 {
		s.list[s.top1] = x
		s.top1++
	} else {
		s.list[s.top2] = x
		s.top2--
	}
	return true
}

func (s *ShareStack) Pop(stackName int, x *int) bool {
	if stackName == 0 {
		if s.top1 == 0 {
			return false
		}
		s.top1--
		*x = s.list[s.top1]
	} else {
		if s.top2 == s.max-1 {
			return false
		}
		s.top2++
		*x = s.list[s.top1]
	}
	return true
}

func (s ShareStack) GetTop(x *int) bool {
	panic("implement me")
}

func (s ShareStack) Destroy() {
	panic("implement me")
}

func (s ShareStack) Print(str, fmts string) {
	if str == "" || str == "0" {
		fmt.Println("stack0:")
		for i := s.top1 - 1; i >= 0; i-- {
			fmt.Printf("%v%v", s.list[i], fmts)
		}
		fmt.Println()
	}
	if str == "" || str == "1" {
		fmt.Println("stack1:")
		for i := s.top2 + 1; i < s.max; i++ {
			fmt.Printf("%v%v", s.list[i], fmts)
		}
		fmt.Println()
	}
}
