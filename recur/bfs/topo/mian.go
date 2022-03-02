package topo

func canFinish(numCourses int, prerequisites [][]int) bool {
	to := make([][]int, numCourses)
	inDeg := make([]int, numCourses)
	for _, prerequisite := range prerequisites {
		a := prerequisite[0]
		b := prerequisite[1]
		to[b] = append(to[b], a)
		inDeg[a]++
	}
	var deque []int
	for i := 0; i < numCourses; i++ {
		if inDeg[i] == 0 {
			deque = append(deque, i)
		}
	}

	var learned []int

	for len(deque) != 0 {
		front := deque[0]
		deque = deque[1:len(deque)]
		learned = append(learned, front)
		for _, a := range to[front] {
			inDeg[a]--
			if inDeg[a] == 0 {
				deque = append(deque, a)
			}
		}
	}
	return len(learned) == numCourses
}
