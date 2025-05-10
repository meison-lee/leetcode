package main

import (
	"fmt"
)

func twoSum(numbers []int, target int) []int {
	
	if len(numbers) == 2 {
		return []int{1,2}
	}

	left, right := 0, len(numbers)-1
	var add int
	for left < right {
		add = numbers[left] + numbers[right]
		if target > add {
			left++
		}else if target < add {
			right--
		}else {
			return []int{left+1, right+1}
		}
		
	}
   	return nil 
}

func main(){
	result := twoSum([]int{1,2,3,4}, 3)
	fmt.Println(result)
}
