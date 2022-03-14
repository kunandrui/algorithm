package main

import (
	"fmt"
	"math/rand"
)

func main() {
	nums := []int{5, 42, 31, 23, 10, 7, 6, 8, 9, 10}
	pivot := partition(nums, 0, 9)
	fmt.Println(nums)
	fmt.Println(pivot)
}

func partition(nums []int, l, r int) int {
	pivot := l + int(rand.Float64()*float64(r-l+1))
	fmt.Println(pivot)
	pivotVal := nums[pivot]
	for l <= r {
		for nums[l] < pivotVal {
			l++
		}
		for nums[r] > pivotVal {
			r--
		}
		if l == r {
			break
		}
		if l < r {
			tmp := nums[l]
			nums[l] = nums[r]
			nums[r] = tmp
			l++
			r--
		}
	}
	return r
}
