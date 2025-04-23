package main

import (
	"fmt"
	"leetcode/lib/sort"
)

func containDuplicate(nums []int) bool {

	nums = sort.MergeSort(nums)
	for i := range nums[:len(nums)-1] {
		if nums[i] == nums[i+1] {
			return true
		}
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
