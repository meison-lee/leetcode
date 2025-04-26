package main

import (
	"fmt"
	"leetcode/lib/sort"
)

// If we use merge sort, it will cost
// time  O(nlogn)
// space O(n) -> from merge sort
// If we use hashSet to implement, we would only need
// time  O(n)
// space O(n)
func containDuplicate(nums []int) bool {

	hash := make(map[int]bool)
	for _, num := range nums {
		if _, ok := hash[num]; !ok {
			hash[num] = true
			continue
		}
		return true
	}

	return false
}

func main() {
	nums := []int{1, 2, 3, 1}
	output := containDuplicate(nums)

	if output {
		fmt.Println("it contain duplicate numbers")
	} else {
		fmt.Println("all unique")
	}

	count := sort.MergeSort([]int{10, 30, 23, 453, 2, 123})
	fmt.Println(count)
}
