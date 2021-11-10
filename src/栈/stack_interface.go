package 栈

type Stack interface {
	InitStack(max int)
	StackEmpty() bool
	Push(x int) bool
	Pop(x *int) bool
	GetTop(x *int) bool
	Destroy()
	Print(s, fmts string)
}
