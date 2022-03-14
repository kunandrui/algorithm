package search

func search(nums []int, target int) int {
	n := len(nums)
	l := 0
	r := n - 1

	for l <= r {
		mid := (l + r) >> 1
		if nums[mid] == target {
			return mid
		}
		//
		if nums[0] <= nums[mid] {
			// 0 ~ mid 有序
			if nums[0] <= target && target < nums[mid] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			// mid ~ n 有序
			if nums[mid] < target && target <= nums[n-1] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}
	return -1
}
