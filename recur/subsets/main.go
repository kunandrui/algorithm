package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 1, 3}
	permuteUnique(nums)
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
