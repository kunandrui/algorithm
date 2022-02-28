package canFinish

/*
 * 拓扑排序
 */

func canFinish(numCourses int, prerequisites [][]int) bool {
	to := make([][]int, numCourses)
	inDeg := make([]int, numCourses)
	for _, pre := range prerequisites {
		a := pre[0]
		b := pre[1]
		to[b] = append(to[b], a)
		inDeg[a]++
	}

	var deque []int
	var lessons []int
	for i := 0; i < numCourses; i++ {
		if inDeg[i] == 0 {
			deque = append(deque, i)
		}
	}
	for len(lessons) != 0 {
		front := deque[0]
		deque = deque[1:len(deque)]
		lessons = append(lessons, front)

		for _, a := range to[front] {
			inDeg[a]--
			if inDeg[a] == 0 {
				deque = append(deque, a)
			}
		}
	}
	return len(lessons) == numCourses
}
