package main

import (
	"fmt"
)

func productExceptSelf(nums []int) (result []int) {
	result = make([]int, len(nums))
	result[0] = 1
	for i := 1; i < len(nums); i++ {
		result[i] = result[i-1] * nums[i-1]	
	}
	
	right := make([]int, len(nums))
        right[len(nums)-1] = 1 
	for i := len(nums)-2; i >= 0; i-- {
		right[i] = right[i+1] * nums[i+1]
	}

	for i := 0; i < len(nums)-1; i++{
		result[i] *=  right[i]
	}

	return result
}

func main() {
	nums := []int{1, 2, 3, 4}
	result := productExceptSelf(nums)
	fmt.Println(result)
}
