package findMin

func findMin(nums []int) int {
	n := len(nums)
	left := 0
	right := n - 1

	for left < right {
		mid := (left + right) >> 1
		if nums[mid] < nums[n-1] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return right
}
