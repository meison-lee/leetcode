package main

import "fmt"

// use map[difference]index
func twoSum(nums []int, target int) []int {
	hash := make(map[int]int, len(nums))
	for index, num := range nums {
		if preIndex, ok := hash[num]; ok {
			return []int{preIndex, index}
		}
		hash[target-num] = index
	}
	return nil
}

func main() {
	result := twoSum([]int{3, 2, 4}, 6)
	fmt.Println(result)
}
