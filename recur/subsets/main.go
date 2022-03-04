package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{10, 1, 10, 9, 6, 1, 9, 5, 9, 10, 7, 8, 5, 2, 10, 8}
	if canPartitionKSubsets(nums, 11) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}

func subsets(nums []int) [][]int {
	var sets [][]int
	var set []int

	var recur func([]int, int)

	recur = func(nums []int, i int) {
		if i == len(nums) {
			sets = append(sets, set)
			//sets = append(sets, append([]int(nil), set...))
			return
		}

		recur(nums, i+1)

		set = append(set, nums[i])
		recur(nums, i+1)
		set = set[:len(set)-1]
	}

	recur(nums, 0)
	fmt.Printf("%v", sets)
	return sets
}

func permuteUnique(nums []int) [][]int {
	sort.Slice(nums, func(i int, j int) bool {
		return nums[i] < nums[j]
	})

	n := len(nums)
	var ans [][]int
	var tmp []int
	used := make([]bool, n)
	var recur func(int)
	recur = func(pos int) {
		if pos == n {
			ans = append(ans, append([]int(nil), tmp...))
			return
		}
		for i := 0; i < n; i++ {

			if !used[i] {
				if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
					continue
				}

				used[i] = true
				tmp = append(tmp, nums[i])
				recur(pos + 1)
				tmp = tmp[:len(tmp)-1]
				used[i] = false
			}
		}
	}
	recur(0)
	fmt.Println(ans)
	return ans
}

func subsets2(nums []int) [][]int {
	n := len(nums)
	var ans [][]int

	var set []int

	var dfs func(int)
	dfs = func(i int) {
		if i == n {
			ans = append(ans, append([]int(nil), set...))
			return
		}
		//决策
		dfs(i + 1)

		set = append(set, nums[i])
		dfs(i + 1)
		set = set[:len(set)-1]
	}
	dfs(0)
	return ans
}

func pailie(nums []int) [][]int {
	n := len(nums)
	var ans [][]int

	var set []int
	used := make([]bool, n)

	var dfs func(int)
	dfs = func(i int) {
		if i == n {
			//边界
			ans = append(ans, append([]int(nil), set...))
			return
		}
		// 决策
		for k := 0; k < n; k++ {
			if !used[k] {
				used[k] = true
				set = append(set, nums[i])
				dfs(i + 1)
				set = set[:len(set)-1]
				used[k] = false
			}
		}
	}
	dfs(0)
	return ans
}

// 分成k个相等的子集
func canPartitionKSubsets(nums []int, k int) bool {
	n := len(nums)
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%k != 0 {
		return false
	}
	target := sum / k
	buckets := make([]int, k)
	var dfs func(int) bool
	dfs = func(i int) bool {
		if i == n {
			for j := 0; j < k; j++ {
				if buckets[j] != target {
					return false
				}
			}
			return true
		}
		// 决策
		for j := 0; j < k; j++ {
			// 装不下 不选
			if buckets[j]+nums[i] > target {
				continue
			}
			buckets[j] += nums[i]
			if dfs(i + 1) {
				return true
			}
			buckets[j] -= nums[i]
		}
		return false
	}
	return dfs(0)
}
