package main

import "sort"

func main() {

}

// intervals = [[1,3],[2,6],[8,10],[15,18]]

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		} else {
			return intervals[i][0] < intervals[j][0]
		}
	})

	var ans [][]int
	n := len(intervals)

	leftBound := intervals[0][0]
	rightBound := intervals[0][1]

	for i := 1; i < n; i++ {
		if intervals[i][0] > rightBound {
			ans = append(ans, []int{leftBound, rightBound})
			leftBound = intervals[i][0]
			rightBound = intervals[i][1]

		} else {
			rightBound = max(rightBound, intervals[i][1])
		}
	}
	ans = append(ans, []int{leftBound, rightBound})

	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
