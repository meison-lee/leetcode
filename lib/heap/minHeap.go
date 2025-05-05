package heap

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	val := old[n-1]
	*h = old[0 : n-1]
	return val
}

func (h *MinHeap) HeapToInt() []int {
	ans := make([]int, h.Len())
	for i := h.Len() - 1; i >= 0; i-- {
		ans[i] = h.Pop().(int)
	}
	return ans
}
