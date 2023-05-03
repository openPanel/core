package router

import (
	"container/heap"
)

type PQNode struct {
	id       string
	distance int
}

var _ heap.Interface = (*PriorityQueue)(nil)

// PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*PQNode

func (pq *PriorityQueue) Push(x any) {
	item := x.(*PQNode)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) Len() int {
	return len(*pq)
}

func (pq *PriorityQueue) Less(i, j int) bool {
	return (*pq)[i].distance < (*pq)[j].distance
}

func (pq *PriorityQueue) Swap(i, j int) {
	(*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i]
}
