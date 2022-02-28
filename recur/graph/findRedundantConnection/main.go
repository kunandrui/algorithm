package main

//冗余连接

func main() {

}

func findRedundantConnection(edges [][]int) []int {
	var ans []int

	n := 0
	for _, edge := range edges {
		x := edge[0]
		y := edge[1]
		n = max(n, max(x, y))
	}
	n++
	to := make([][]int, n)

	visited := make([]bool, n)
	var hasCycle bool

	var dfs func(x int, fa int)
	dfs = func(x int, fa int) {
		visited[x] = true

		for _, y := range to[x] {
			if y == fa {
				continue
			}
			if !visited[y] {
				dfs(y, x)
			} else {
				//有环
				hasCycle = true
			}

		}
	}

	for _, edge := range edges {
		x := edge[0]
		y := edge[1]
		to[x] = append(to[x], y)
		to[y] = append(to[y], x)
		for i := 0; i < n; i++ {
			visited[i] = false
		}
		hasCycle = false
		dfs(x, 0)
		if hasCycle {
			return edge
		}
	}

	return ans
}

func max(x int, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
