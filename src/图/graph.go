package 图

const MAXPOINT = 100

type AdjacencyMatrix struct {
	V [MAXPOINT]interface{}
	E [MAXPOINT][MAXPOINT]interface{}
}

func testAdjacencyMatrix() {
	var test AdjacencyMatrix
	test.E[1][1] = 10
	test.V[0] = 123
}

type lTableEdge struct {
	index, nextIndex int
}

type lTablePoint struct {
	v    interface{}
	next *lTableEdge
}

type LeadingTable struct {
	V []lTablePoint
}

func testLeadingTable() {
	var test LeadingTable
	test.V[1].v = 20
	test.V[1].next = new(lTableEdge)
	test.V[1].next.index = 3
	test.V[1].next.nextIndex = -1
}

type crossLinkEdge struct {
	tail, head   int
	hLink, tLink *crossLinkEdge
}

type crossLinkPoint struct {
	v         interface{}
	fin, fout *crossLinkEdge
}

type CrossLinkedList struct {
	V []crossLinkPoint
}

func testCrossLinkedList() {
	var test CrossLinkedList
	test.V[0].v = 10
	p := new(crossLinkEdge)
	test.V[0].fin = p
	p.tail = 0
	p.head = 1
	test.V[0].fout = new(crossLinkEdge)
}
