package threeSum

import "sort"

func threeSum(nums []int) [][]int {
	var ans [][]int

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		target := -nums[i]
		twoSumPairs := twoSum(nums, i+1, target)
		for _, twoSumPair := range twoSumPairs {
			ans = append(ans, []int{nums[i], twoSumPair[0], twoSumPair[1]})
		}
	}
	return ans
}

func twoSum(nums []int, start int, target int) [][]int {
	var ans [][]int

	l := start
	r := len(nums) - 1

	for ; l < r; l++ {
		if l > start && nums[l] == nums[l-1] {
			continue
		}
		for l < r && nums[l]+nums[r] > target {
			r--
		}
		if l < r && nums[l]+nums[r] == target {
			ans = append(ans, []int{nums[l], nums[r]})
		}
	}
	return ans
}
