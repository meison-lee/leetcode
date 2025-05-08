package main

import (
	"fmt"
)

func topKFrequent(nums []int, k int) []int {
	hash := make(map[int]int)
	for _, v := range nums {
		hash[v]++
	}

	buckets := make([][]int, len(nums)+1)
	for i, v := range hash {
		buckets[v] = append(buckets[v], i)
	}

	result := []int{}
	for i := len(buckets) - 1; i >= 0; i-- {
		for _, v := range buckets[i] {
			result = append(result, v)
			if len(result) == k {
				return result
			}
		}
	}
	return result
}

func main() {
	result := topKFrequent([]int{1, 1, 1, 7, 7, 5, 4, 4, 4}, 3)
	fmt.Println(result)
}
