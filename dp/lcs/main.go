package main

import "fmt"

func main() {
	text1 := "abcde"
	text2 := "ace"
	ans := longestCommonSubsequence(text1, text2)
	fmt.Println(ans)
}

func longestCommonSubsequence(text1 string, text2 string) int {
	m := len(text1)
	n := len(text2)
	text1 = " " + text1
	text2 = " " + text2

	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}

	ans := 0
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if text1[i] == text2[j] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = getMax(dp[i-1][j], dp[i][j-1])
			}

			if ans < dp[i][j] {
				ans = dp[i][j]
			}
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
