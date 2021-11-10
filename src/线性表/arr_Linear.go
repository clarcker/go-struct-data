package 线性表

import "fmt"

type List struct {
	arr    []int
	nowLen int
	maxLen int
}

func (l *List) DeleteAllX(x int) {
	if l.nowLen == 0 {
		return
	}
	i := 0
	for j := 0; j < l.nowLen; j++ {
		if l.arr[j] == x {
			i++
		} else {
			l.arr[j-i] = l.arr[j]
		}
	}
	l.nowLen -= i
}

func (l *List) DeleteST(s, t int) {
	if s > t {
		fmt.Println("s:t error", s, t)
		return
	}
	k := 0
	for j := 0; j < l.nowLen; j++ {
		if l.arr[j] >= s && l.arr[j] <= t {
			k++
		} else {
			l.arr[j-k] = l.arr[j]
		}
	}
	l.nowLen -= k
}

func (l *List) DeleteAllSame() {
	k := 0
	for j := 1; j < l.nowLen; j++ {
		if l.arr[j] == l.arr[j-1] {
			k++
		} else {
			l.arr[j-k] = l.arr[j]
		}
	}
	l.nowLen -= k
}

func (l *List) Init(lens int) {
	l.maxLen = lens
	l.nowLen = 0
	l.arr = make([]int, lens)
}

func (l *List) InsertEnd(e int) bool {
	if l.nowLen >= l.maxLen {
		return false
	}
	l.arr[l.nowLen] = e
	l.nowLen++
	return true
}

func (l *List) InsertLocate(index int, e int) bool {
	if l.nowLen >= l.maxLen || index < 0 || index > l.nowLen {
		return false
	}
	i := l.nowLen
	for ; i > index; i-- {
		l.arr[i] = l.arr[i-1]
	}
	l.arr[i] = e
	l.nowLen++
	return true
}

func (l List) LocateElem(e int) int {
	for i, v := range l.arr {
		if v == e {
			return i
		}
	}
	return -1
}

func (l List) GetElem(index int) (int, bool) {
	if index < 0 || index >= l.nowLen {
		return -1, false
	}
	return l.arr[index], true
}

func (l *List) DeleteElem(e int) bool {
	index := l.LocateElem(e)
	if index != -1 {
		for i := index; i < l.nowLen-1; i++ {
			l.arr[i] = l.arr[i+1]
		}
		l.nowLen--
		return true
	} else {
		return false
	}
}

func (l *List) DeleteIndex(index int) bool {
	if index < 0 || index >= l.nowLen {
		return false
	}
	for i := index; i < l.nowLen-1; i++ {
		l.arr[i] = l.arr[i+1]
	}
	l.nowLen--
	return true
}

func (l List) Empty() bool {
	return l.nowLen == 0
}

func (l List) Length() int {
	return l.nowLen
}

func (l List) Print(a string) {
	for i := 0; i < l.nowLen; i++ {
		fmt.Printf("%v%s", l.arr[i], a)
	}
	fmt.Println()
}

func (l *List) Destroy() {
	l.nowLen = 0
	l.arr = nil
	l.maxLen = 0
}
func (l *List) FindX(x int) {
	i, j := 0, l.nowLen-1
	for i <= j {
		if l.arr[(i+j)/2] > x {
			j = (i+j)/2 - 1
		} else if l.arr[(i+j)/2] < x {
			i = (i+j)/2 + 1
		} else {
			if i != l.nowLen {
				l.arr[(i+j)/2], l.arr[(i+j)/2+1] = l.arr[(i+j)/2+1], l.arr[(i+j)/2]
				return
			}
		}
	}
	fmt.Println(l.arr[j])
	if i > j {
		l.InsertLocate(j+1, x)
	}

}

func AddArrLinear(l1, l2 List) List {
	var ret List
	ret.Init(l1.maxLen + l2.maxLen)
	i, j, k := 0, 0, 0
	for ; i < l1.nowLen || j < l2.nowLen; k++ {
		if i < l1.nowLen && j < l2.nowLen {
			if l1.arr[i] < l2.arr[j] {
				ret.arr[k] = l1.arr[i]
				i++
			} else {
				ret.arr[k] = l2.arr[j]
				j++
			}
		} else if i < l1.nowLen && j >= l2.nowLen {
			ret.arr[k] = l1.arr[i]
			i++
		} else if j < l2.nowLen && i >= l1.nowLen {
			ret.arr[k] = l2.arr[j]
			j++
		}
		ret.nowLen++
	}
	return ret
}

func Marjority(l List) int {
	m, count := l.arr[0], 0
	for _, data := range l.arr {
		if data == m {
			count++
		} else {
			if count > 0 {
				count--
			} else {
				m = data
				count = 1
			}
		}
	}
	count = 0
	for _, data := range l.arr {
		if data == m {
			count++
		}
	}
	if count > l.nowLen/2 {
		return m
	} else {
		return -1
	}
}
