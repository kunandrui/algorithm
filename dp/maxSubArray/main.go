package main

func main() {

}

func maxSubArray(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	dp[0] = nums[0]

	ans := -10000
	for i := 1; i < n; i++ {
		dp[i] = getMax(dp[i-1]+nums[i], nums[i])
		if ans < dp[i] {
			ans = dp[i]
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
