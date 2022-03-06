package main

import (
	"math/rand"
)

func main() {

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
	pivot := l + int(rand.Float64()*float64(r-l+1))
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
