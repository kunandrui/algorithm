package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 1}
	fmt.Println(rob(nums))
}

func rob(nums []int) int {
	n := len(nums)
	nums = append([]int{0}, nums...)
	//f[i,0] <- max(f[i-1,1], f[i-1,0] + nums[i])
	//f[i,1] <- f[i-1, 0]
	INF := int(-1e9)
	f := make([][2]int, n+1)
	for i := 0; i <= n; i++ {
		for j := 0; j < 2; j++ {
			f[i][j] = INF
		}
	}

	f[0][0] = 0
	for i := 1; i <= n; i++ {
		f[i][1] = f[i-1][0] + nums[i]
		f[i][0] = max(f[i-1][1], f[i-1][0])
	}

	return max(f[n][0], f[n][1])
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
