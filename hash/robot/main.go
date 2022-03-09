package robot

import "strconv"

func main() {
	//set := make(map[string]struct{})
}

func robotSim(commands []int, obstacles [][]int) int {
	set := make(map[string]struct{})
	for _, obstacle := range obstacles {
		x := obstacle[0]
		y := obstacle[1]
		set[hash(x, y)] = struct{}{}
	}

	dx := []int{0, 1, 0, -1}
	dy := []int{1, 0, -1, 0}

	x := 0
	y := 0
	ans := 0
	dir := 0
	for _, command := range commands {
		if command == -1 {
			dir = (dir + 1) % 4
		} else if command == -2 {
			dir = (dir + 3) % 4
		} else {
			for i := 0; i < command; i++ {
				nx := x + dx[dir]
				ny := y + dy[dir]
				if _, exists := set[hash(nx, ny)]; exists {
					break
				}
				x = nx
				y = ny
			}
			ans = max(ans, x*x+y*y)
		}
	}
	return ans
}

func hash(x, y int) string {
	return strconv.Itoa(x) + "," + strconv.Itoa(y)
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
