package maxProduct

func maxProduct(nums []int) int {
	n := len(nums)
	dpMax := make([]int, n)
	dpMin := make([]int, n)
	dpMax[0] = nums[0]
	dpMin[0] = nums[0]

	for i := 1; i < n; i++ {
		dpMax[i] = getMax(getMax(dpMax[i-1]*nums[i], dpMin[i-1]*nums[i]), nums[i])
		dpMin[i] = getMin(getMin(dpMax[i-1]*nums[i], dpMin[i-1]*nums[i]), nums[i])
	}
	ans := dpMax[0]
	for i := 0; i < n; i++ {
		if ans < dpMax[i] {
			ans = dpMax[i]
		}
	}
	return ans
}

func getMax(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func getMin(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}
