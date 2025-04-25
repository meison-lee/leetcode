package sort

func CountingSortByDigit(nums []int, place int) []int {

	count := make([]int, 10)
	for _, num := range nums {
		digit := GetDigit(num, place)
		count[digit]++
	}

	for i := 1; i < len(count); i++ {
		count[i] += count[i-1]
	}

	output := make([]int, len(nums))
	for i := len(nums) - 1; i >= 0; i-- {
		digit := GetDigit(nums[i], place)
		output[count[digit]-1] = nums[i]
		count[digit]--
	}

	return output
}

func GetDigit(num, place int) int {
	return (num / place) % 10
}

func GetMax(nums []int) int {

	largest := 0
	for _, num := range nums {
		if num > largest {
			largest = num
		}
	}
	return largest
}

func RadixSort(nums []int) []int {

	max := GetMax(nums)
	for place := 1; max/place > 0; place *= 10 {
		nums = CountingSortByDigit(nums, place)
	}
	return nums
}
