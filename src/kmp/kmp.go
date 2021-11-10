package kmp

import "fmt"

type Strs struct {
	s string
}

type Strser interface {
	StrAssign(t string) bool
	StrCopy(t string) bool
	StrEmpty() bool
	StrCompare(t Strs) int
	StrLen() int
	SubString(s []string, pos, len int)
	Index(t Strs) int
	ClearString()
	DestroyString()
}

func (s *Strs) StrAssign(t string) bool {
	s.s = t
	return true
}

func (s *Strs) StrCopy(t string) bool {
	s.s = fmt.Sprintf("%v%v", s.s, t)
	return true
}

func (s Strs) StrEmpty() bool {
	return len(s.s) == 0
}

func nextArray(t string) []int {
	ret := make([]int, len(t))
	ret[0] = -1
	k := -1
	for i := 0; i < len(t); {
		if k == -1 || t[k] == t[i] {
			k++
			ret[i] = k
			i++
		} else {
			k = ret[k]
		}
	}
	return ret
}

func (s Strs) Index(t string) int { //kmp
	next := nextArray(t)
	i, j := 0, 0
	for i = 0; i < len(s.s) && j < len(t); {
		if s.s[i] == t[i] || j == -1 {
			i++
		} else {
			j = next[j]
		}
	}
	if j == len(t) {
		return i - j
	} else {
		return -1
	}
}

func (s Strs) ClearString() {

	panic("implement me")
}

func (s Strs) DestroyString() {
	panic("implement me")
}
