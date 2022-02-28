package main

func main() {
	bank := []string{"AACCGGTA", "AACCGCTA", "AAACGGTA"}
	minMutation("AACCGGTT", "AAACGGTA", bank)
}

/*
"ACGT"
start: "AACCGGTT"
end:   "AACCGGTA"
bank: ["AACCGGTA"]
*/
type void struct{}

var member void

func minMutation(start string, end string, bank []string) int {
	seq := []byte("ACGT")
	hash := make(map[string]int)
	set := make(map[string]void)

	for _, str := range bank {
		set[str] = member
	}

	if _, exist := set[end]; !exist {
		return -1
	}

	var deque []string
	deque = append(deque, start)
	hash[start] = 0
	for len(deque) != 0 {
		front := deque[0]
		deque = deque[1:len(deque)]
		for i := 0; i < len(front); i++ {
			for j := 0; j < len(seq); j++ {
				if front[i] != seq[j] {
					tmp := []byte(front)
					tmp[i] = seq[j]
					ns := string(tmp)

					if _, exist := set[ns]; !exist {
						continue
					}

					if _, exist := hash[ns]; exist {
						continue
					}

					deque = append(deque, ns)
					hash[ns] = hash[front] + 1

					if ns == end {
						return hash[ns]
					}
				}
			}
		}
	}

	return -1
}

/**
 * dfs
 */
func longestIncreasingPath1(matrix [][]int) int {
	m := len(matrix)
	n := len(matrix[0])
	dist := make([][]int, m)
	for i := 0; i < m; i++ {
		dist[i] = make([]int, n)
	}

	dx := []int{-1, 0, 1, 0}
	dy := []int{0, 1, 0, -1}

	var dfs func(x int, y int) int
	dfs = func(x int, y int) int {
		if dist[x][y] != 0 {
			return dist[x][y]
		}
		dist[x][y] = 1
		for k := 0; k < 4; k++ {
			nx := x + dx[k]
			ny := y + dy[k]

			if nx >= 0 && nx < m && ny >= 0 && ny < n && matrix[nx][ny] > matrix[x][y] {
				dist[x][y] = max(dist[x][y], dfs(nx, ny)+1)
			}
		}
		return dist[x][y]
	}
	ans := 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			ans = max(ans, dfs(i, j))
		}
	}

	return ans
}

/**
 * bfs
 */
func longestIncreasingPath2(matrix [][]int) int {
	m := len(matrix)
	n := len(matrix[0])

	dist := make([]int, m*n)
	to := make([][]int, m*n)
	inDeg := make([]int, m*n)

	ans := 0

	dx := []int{-1, 0, 1, 0}
	dy := []int{0, 1, 0, -1}

	var num func(x int, y int) int
	num = func(x int, y int) int {
		return x*n + y
	}

	for x := 0; x < m; x++ {
		for y := 0; y < n; y++ {
			for k := 0; k < 4; k++ {
				nx := x + dx[k]
				ny := y + dy[k]
				if nx >= 0 && nx < m && ny >= 0 && ny < n && matrix[nx][ny] > matrix[x][y] {
					to[num(x, y)] = append(to[num(x, y)], num(nx, ny))
					inDeg[num(nx, ny)]++
				}
			}
		}
	}

	var deque []int
	for i := 0; i < m*n; i++ {
		if inDeg[i] == 0 {
			dist[i] = 1
			deque = append(deque, i)
		}
	}
	for len(deque) != 0 {
		x := deque[0]
		deque = deque[1:len(deque)]
		for _, y := range to[x] {
			inDeg[y]--
			dist[y] = max(dist[y], dist[x]+1)
			if inDeg[y] == 0 {
				deque = append(deque, y)
			}
		}
	}

	for i := 0; i < m*n; i++ {
		ans = max(ans, dist[i])
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
