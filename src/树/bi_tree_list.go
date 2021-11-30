package 树

import (
	"fmt"
	"go-struct-data/栈"
	"math"
)

type BiTreeList struct {
	tree []int
	k    int
}

func (b *BiTreeList) InitTree(n int) {
	b.k = n
	max := int(math.Pow(2, float64(n)))
	b.tree = make([]int, max)
	b.tree[0] = max
}

func (b *BiTreeList) InsertTreeByIndex(index, x int) bool {
	if index > b.tree[0] {
		return false
	}
	b.tree[index] = x
	return true
}

func (b *BiTreeList) CopyArrayToTree(arr []int) {
	len := len(arr)
	n := int(math.Ceil(math.Log2(float64(len + 1))))
	b.InitTree(n)
	for i := 0; i < len; i++ {
		b.tree[i+1] = arr[i]
	}
}

func (b BiTreeList) PreVisit() {
	var stack 栈.ListStack
	stack.InitStack(b.tree[0] - 1)
	j := 1
	for !stack.StackEmpty() || j < b.tree[0] {
		if j < b.tree[0] && b.tree[j] != 0 {
			fmt.Printf("%v ", b.tree[j])
			stack.Push(j)
			j *= 2
		} else {
			stack.Pop(&j)
			j = j*2 + 1
		}
	}
	fmt.Println()
}

func (b BiTreeList) MidVisit() {
	var stack 栈.ListStack
	stack.InitStack(b.tree[0] - 1)
	j := 1
	for !stack.StackEmpty() || b.tree[0] > j {
		if j < b.tree[0] && b.tree[j] != 0 {
			stack.Push(j)
			j *= 2
		} else {
			stack.Pop(&j)
			if j < b.tree[0] && b.tree[j] != 0 {
				fmt.Printf("%v ", b.tree[j])
			}
			j = j*2 + 1
		}
	}
	fmt.Println()
}

func (b BiTreeList) PostVisit() {
	visit := make(map[int]int)
	var stack 栈.ListStack
	stack.InitStack(b.tree[0] - 1)
	j := 1
	visit[j] = 1
	stack.Push(j)
	for !stack.StackEmpty() || j != -1 {
		if j < b.tree[0] && b.tree[j] != 0 {
			if _, ok := visit[j*2]; ok {

			} else {
				if j != 1 {
					stack.Push(j)
				}
				j *= 2
				visit[j] = 1
				continue
			}
			if _, ok := visit[j*2+1]; ok {
				fmt.Printf("%v ", b.tree[j])
				if !stack.Pop(&j) {
					j = -1
				}
			} else {
				stack.Push(j)
				j = j*2 + 1
				visit[j] = 1
				continue
			}
		} else {
			stack.Pop(&j)
		}
	}
	fmt.Println()
}

func (b BiTreeList) LevelVisit() {
	for i := 1; i < b.tree[0]; i++ {
		if b.tree[i] != 0 {
			fmt.Printf("%v ", b.tree[i])
		}
	}
	fmt.Println()
}

func (b BiTreeList) TreePrint(s string) {
	fmt.Println(s)
	fmt.Println(b.k, b.tree[0])
	f := func(n int) {
		for i := 0; i < n; i++ {
			fmt.Printf("-")
		}
	}
	j := 1
	for i := 0; i < b.k; i++ {
		//qian
		qian := int(math.Pow(2, float64(b.k-i-1))) - 1
		if i == 0 {
			f((b.tree[0] - 1) / 2)
		} else if i == b.k-1 {

		} else {
			f(qian)
		}
		//zhong
		zhong := int(math.Pow(2, float64(b.k-i))) - 1
		for ; j <= int(math.Pow(2, float64(i+1)))-1; j++ {
			fmt.Printf("%v", b.tree[j])
			if j != int(math.Pow(2, float64(i+1)))-1 {
				f(zhong)
			}
		}
		if i == 0 {
			f((b.tree[0] - 1) / 2)
		} else if i == b.k-1 {

		} else {
			f(qian)
		}
		fmt.Println()
	}
}
