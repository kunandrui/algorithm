package sssp

// 1. 初始化 dist[1] = 0
// 2. 找出未被标记的, 最小的dist[x], 标记x
// 3. 扫描x的出边, 若dist[y] > dist[x] + z, 则更新dist[y]
// 4. 重复2~3

func networkDelayTime02(times [][]int, n int, k int) int {
	to := make([][]int, n+1)
	edge := make([][]int, n+1)
	for _, time := range times {
		x := time[0]
		y := time[1]
		z := time[2]
		to[x] = append(to[x], y)
		edge[x] = append(edge[x], z)
	}

	dist := make([]int, n+1)
	for i := 0; i <= n; i++ {
		dist[i] = int(1e9)
	}
	dist[k] = 0

	expand := make([]bool, n+1)
	expand[k] = true
	for round := 1; round <= n-1; round++ {
		min := int(1e9)
		x := 0
		for i := 1; i <= n; i++ {
			if !expand[i] && dist[i] < min {
				min = dist[i]
				x = i
			}
		}
		expand[x] = true
		for i := 0; i < len(to[x]); i++ {
			y := to[x][i]
			z := edge[x][i]
			if dist[x]+z < dist[y] {
				dist[y] = dist[x] + z
			}
		}
	}
	return dist[1]
}
