package main

import (
	"fmt"
	"go-struct-data/栈"
	"go-struct-data/线性表"
	"go-struct-data/队列"
	"math/rand"
	"time"
)

func testList() {
	var test1 = 线性表.List{}
	for j := 0; j < 5; j++ {
		test1.Init(20)
		rand.Seed(time.Now().UnixNano())
		for i := 0; i < 10; i++ {
			e := rand.Intn(2)
			test1.InsertEnd(i*2 + e)
		}
		x := rand.Intn(25)
		fmt.Println(x)
		test1.Print(" ")
		test1.FindX(x)
		test1.Print(" ")
		test1.Destroy()
	}
}

func testDoubleNodeLinear() {
	/*
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
	*/
	var test 线性表.DoubleNodeLinear
	var e int
	for i := 0; i < 20; i++ {
		rand.Seed(time.Now().UnixNano())
		e = rand.Intn(20)
		if i%2 == 0 {
			test.InsertHead(e)
		} else {
			test.InsertEnd(e)
		}
		fmt.Println(test.Length())
	}
	test.Print(" ")
	for i := 0; i < 50; i++ {
		rand.Seed(time.Now().UnixNano())
		e = rand.Intn(25)
		fmt.Println(e)
		if e%2 == 0 {
			index := test.DeleteElem(e)
			if index {
				fmt.Printf("DeleteElem:rand:%v\n", e)
			} else {
				fmt.Printf("DeleteElem:rand:%v delete fail\n", e)
			}
		} else {
			ok := test.DeleteIndex(e)
			if ok {
				fmt.Printf("DeleteIndex:rand:%v\n", e)
			} else {
				fmt.Printf("DeleteIndex:rand:%v  delete fail\n", e)
			}
		}
		test.Print(" ")
	}
}

func testNodeList() {
	var test 线性表.NodeList
	//rand.Seed(time.Now().UnixNano())
	for i := 0; i < 19; i++ {
		//e := rand.Intn(10)
		//test.InsertSort(e)
		//test.Print(" ")
		test.InsertEnd(i)
	}
	test.Print(" ")
	test.P39T25()
	test.Print(" ")
}

func testNodeStack() {
	var test1 栈.NodeStack
	test1.InitStack(10)
	rand.Seed(time.Now().UnixNano())
	if test1.StackEmpty() {
		for i := 0; i < 12; i++ {
			x := rand.Intn(20)
			if test1.Push(x) {
				test1.Print(fmt.Sprint("put int:", x), " ")
			} else {
				var top int
				test1.GetTop(&top)
				test1.Print(fmt.Sprint("put fail top:", top), " ")
			}
		}
	}
	if !test1.StackEmpty() {
		x := rand.Intn(20)
		fmt.Println(x)
		for i := 0; i < x; i++ {
			var e int
			if test1.Pop(&e) {
				test1.Print(fmt.Sprint("pop int:", e), " ")
			} else {
				fmt.Println(test1.StackEmpty())
			}
		}
	}
}

func testShareStack() {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(40)
	var test 栈.ShareStack
	max := rand.Intn(20)
	test.InitStack(max)
	fmt.Printf("max:%v n:%v\n", max, n)
	for i := 0; i < n; i++ {
		r := rand.Intn(2)
		op := rand.Intn(2)
		e := rand.Intn(100)
		if op == 0 {
			if test.Push(r, e) {
				fmt.Printf("push: stack:%v val:%v\n", r, e)
			} else {
				fmt.Printf("push fail: stack:%v val:%v\n", r, e)
			}
		} else {
			var x int
			if test.Pop(r, &x) {
				fmt.Printf("pop: stack:%v val:%v\n", r, e)
			} else {
				fmt.Printf("pop fail: stack:%v val:%v\n", r, e)
			}
		}
		test.Print("", " ")
	}
}

func main() {
	fmt.Println("hello world!")
	//testList()
	//testNodeList()
	testKMP()
}

func testKMP() {
	s := "hello world"
	t := s
	t = "two"
	fmt.Printf("%T %v\n", s[:], s[:])
	fmt.Printf("%T %v\n", t[:], t[:])

}
func testNodeQueue() {
	var test 队列.NodeQueue
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(30)
	for i := 0; i < n; i++ {
		x := rand.Intn(100)
		op := rand.Intn(5)
		if op != 0 {
			if test.EnQueue(x) {
				fmt.Printf("enqueue x:%v op:%v\n", x, op)
			} else {
				fmt.Printf("enqueue fail x:%v op:%v\n", x, op)
			}
		} else {
			if test.DeQueue(&x) {
				fmt.Printf("dequeue x:%v op:%v\n", x, op)
			} else {
				fmt.Printf("dequeue fail ,is empty:%v op:%v\n", test.QueueEmpty(), op)
			}
		}
		test.Print(fmt.Sprintf("i:=%v", i), " ")
	}

}

func testCircularQueue() {
	var test 队列.CircularQueue2
	test.InitQueue(10)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 8; i++ {
		x := rand.Intn(30)
		fmt.Println(x)
		test.EnQueue(x)
	}
	test.Print("en 8", " ")
	for i := 0; i < 5; i++ {
		var x int
		test.DeQueue(&x)
		fmt.Println(x)
	}
	test.Print("De 5", " ")

	for i := 0; i < 15; i++ {
		x := rand.Intn(30)
		if test.EnQueue(x) {
			test.Print(fmt.Sprint(x), " ")
		} else {
			test.Print(fmt.Sprint("fail", x), " ")
		}
	}
	for i := 0; i < 15; i++ {
		var x int
		if test.DeQueue(&x) {
			test.Print(fmt.Sprint(x), " ")
		} else {
			test.Print(fmt.Sprint("fail", x), " ")
		}
	}
	for i := 0; i < 15; i++ {
		x := rand.Intn(30)
		if test.EnQueue(x) {
			test.Print(fmt.Sprint(x), " ")
		} else {
			test.Print(fmt.Sprint("fail", x), " ")
		}
	}
	test.Print("en 15", " ")
}
