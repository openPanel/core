package router

import (
	"container/heap"
)

type pqNode struct {
	id       string
	distance int
}

var _ heap.Interface = (*priorityQueue)(nil)

// priorityQueue implements heap.Interface and holds Items.
type priorityQueue []*pqNode

func (pq *priorityQueue) Push(x any) {
	item := x.(*pqNode)
	*pq = append(*pq, item)
}

func (pq *priorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func (pq *priorityQueue) Len() int {
	return len(*pq)
}

func (pq *priorityQueue) Less(i, j int) bool {
	return (*pq)[i].distance < (*pq)[j].distance
}

func (pq *priorityQueue) Swap(i, j int) {
	(*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i]
}
