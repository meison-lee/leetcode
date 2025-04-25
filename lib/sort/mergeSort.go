package sort

func MergeSort(nums []int) []int {
	if len(nums) > 1 {
		// round up
		half := len(nums) / 2
		left := nums[:half]
		right := nums[half:]

		left = MergeSort(left)
		right = MergeSort(right)

		// compare and combine
		leftPtr, rightPtr := 0, 0
		result := make([]int, len(left)+len(right))
		i := 0

		for leftPtr < len(left) && rightPtr < len(right) {
			if left[leftPtr] < right[rightPtr] {
				result[i] = left[leftPtr]
				leftPtr++
			} else {
				result[i] = right[rightPtr]
				rightPtr++
			}
			i++
		}
		for leftPtr < len(left) {
			result[i] = left[leftPtr]
			leftPtr++
			i++
		}
		for rightPtr < len(right) {
			result[i] = right[rightPtr]
			rightPtr++
			i++
		}
		return result
	}
	return nums
}
