package 队列

type Queue interface {
	InitQueue(max int)
	QueueEmpty() bool
	EnQueue(x int) bool
	DeQueue(x *int) bool
	GetTop(x *int) bool
	Destroy()
	Print(s, fmts string)
}
