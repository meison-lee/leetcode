package main

import (
	"fmt"
	"leetcode/lib/heap"
)

func topKFrequent(nums []int, k int) []int {
	hash := make(map[int]int)
	for _, v := range nums {
		hash[v]++
	}

	fmt.Println(hash)
	h := &heap.MinHeap{}
	// heap.Init(h)

	for _, v := range hash {
		if h.Len() < k {
			h.Push(v)
		} else if v > (*h)[0] {
			h.Pop()
			h.Push(v)
		}
	}

	return h.HeapToInt()
}

func main() {
	result := topKFrequent([]int{1, 1, 1, 2, 2, 3, 4, 4, 4}, 3)
	fmt.Println(result)
}
