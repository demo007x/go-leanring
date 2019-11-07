package container

import "container/heap"

type IntHeap []int // 定义一个堆

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{})  {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n - 1]
	*h = old[0 : n-1]
	return x
}
// 以上IntHeap 实现了堆的结构
func main() {
	h := &IntHeap{2,1,5}
	heap.Init(h)
	heap.Push(h, 3)
	heap.Pop(h)
}