package sssp

// bellman ford 算法
func networkDelayTime(times [][]int, n int, k int) int {
	dist := make([]int, n+1)
	for i := 0; i <= n; i++ {
		dist[i] = int(1e9)
	}
	dist[k] = 0

	for round := 1; round <= n-1; round++ {
		for _, time := range times {
			x := time[0]
			y := time[1]
			z := time[2]

			if dist[x]+z < dist[y] {
				dist[y] = dist[x] + z
			}
		}
	}
	ans := 0
	for i := 1; i <= n; i++ {
		ans = max(ans, dist[i])
	}
	if ans == int(1e9) {
		return -1
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
