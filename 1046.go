package main

import (
	"container/heap"
	"fmt"
)

type PQ []int

func (pq PQ) Len() int { return len(pq) }

func (pq PQ) Less(i, j int) bool { return pq[i] > pq[j] }

func (pq PQ) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }

// Push 往 pq 中放 int
func (pq *PQ) Push(x interface{}) {
	temp := x.(int)
	*pq = append(*pq, temp)
}

// Pop 从 pq 中取出最优先的 int
func (pq *PQ) Pop() interface{} {
	temp := (*pq)[len(*pq)-1]
	*pq = (*pq)[:len(*pq)-1]
	return temp
}

func lastStoneWeight(stones []int) int {
	pq := PQ(stones)
	heap.Init(&pq)
	for pq.Len() >= 2 {
		a, b := heap.Pop(&pq).(int), heap.Pop(&pq).(int)
		if a == b {
			continue
		}
		heap.Push(&pq, a-b)
	}

	if pq.Len() == 0 {
		return 0
	}

	return pq[0]
}

func main() {
	arr := []int{2, 7, 4, 1, 8, 1}
	//arr = []int{1, 3}
	//arr = []int{2, 2}

	fmt.Println(lastStoneWeight(arr))
}
