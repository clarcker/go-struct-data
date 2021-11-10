package 树

import (
	"fmt"
	"math"
)

type BiTreeList struct {
	tree []int
	k    int
}

func (b *BiTreeList) InitTree(n int) {
	b.k = n
	max := int(math.Pow(2, float64(n))) + 1
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

func (b BiTreeList) TreePrint(s string) {
	fmt.Println(s)

}
