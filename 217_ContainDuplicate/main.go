package main

import (
	"fmt"
	"leetcode/lib"
)

func containDuplicate(nums []int) bool {

	for _, num := range nums {
		fmt.Println(num)
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

	count := lib.RadixSort([]int{10, 30, 23, 453, 2, 123})
	fmt.Println(count)
}
