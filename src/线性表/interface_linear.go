package 线性表

type Linear interface {
	Init(lens int)
	InsertEnd(e int) bool
	InsertLocate(index int, e int) bool
	LocateElem(e int) int
	GetElem(index int) (int, bool)
	DeleteElem(e int) bool
	DeleteIndex(index int) bool
	Empty() bool
	Length() int
	Print(s string)
	Destroy()
}
