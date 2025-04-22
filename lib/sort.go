package lib

func CountingSortByDigit(nums []int, place int) []int {

	count := make([]int, 10)
	for _, num := range nums {
		digit := GetDigit(num, place)
		count[digit]++
	}

	for i := 1; i < len(count); i++ {
		count[i] += count[i-1]
	}

	return count
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
