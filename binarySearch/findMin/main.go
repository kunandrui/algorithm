package findMin

func findMin(nums []int) int {
	n := len(nums)
	l := 0
	r := n - 1
	for l < r {
		mid := (l + r) >> 1
		if nums[mid] < nums[n-1] {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return nums[r]
}
