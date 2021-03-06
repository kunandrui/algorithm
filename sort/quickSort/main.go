package main

import (
	"fmt"
	"math/rand"
)

func main() {
	nums := []int{5, 42, 31, 23, 10, 7, 6, 8, 9, 10}
	quickSort(nums, 0, len(nums)-1)
	fmt.Println(nums)
}

func quickSort(nums []int, l, r int) {
	if l >= r {
		return
	}
	pivot := partition(nums, l, r)
	quickSort(nums, l, pivot)
	quickSort(nums, pivot+1, r)
}

func partition(nums []int, l, r int) int {
	pivot := l + rand.Int()%(r-l+1)
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
