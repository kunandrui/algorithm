package main

import "fmt"

func main() {
	nums := []int{10, 9, 9, 7, 6, 5, 4, 3, 2, 1}

	mergeSort(nums, 0, len(nums)-1)
	fmt.Println(nums)
}

func mergeSort(nums []int, l, r int) {
	if l >= r {
		return
	}
	mid := (l + r) >> 1
	mergeSort(nums, l, mid)
	mergeSort(nums, mid+1, r)
	merge(nums, l, mid, r)
}

func merge(nums []int, left, mid, right int) {
	tmp := make([]int, right-left+1)
	i := left
	j := mid + 1
	for k := 0; k < right-left+1; k++ {
		if j > right || (i <= mid && nums[i] < nums[j]) {
			tmp[k] = nums[i]
			i++
		} else {
			tmp[k] = nums[j]
			j++
		}
	}
	for k := 0; k < right-left+1; k++ {
		nums[left+k] = tmp[k]
	}
}
