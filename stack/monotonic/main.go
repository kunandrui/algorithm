package main

import "fmt"

type rect struct {
	width  int
	height int
}

func main() {

	heights := []int{2, 1, 5, 6, 2, 3, 2}

	fmt.Printf("%v", largestRectangleArea(heights))

}

/*
 * 单调栈
 */

func largestRectangleArea(heights []int) int {

	heights = append(heights, 0)

	max := 0

	var stack []rect

	for _, height := range heights {

		accumulatedWidth := 0

		for len(stack) != 0 && height <= stack[len(stack)-1].height {

			accumulatedWidth += stack[len(stack)-1].width

			area := accumulatedWidth * stack[len(stack)-1].height

			if area > max {

				max = area

			}

			stack = stack[:len(stack)-1]

		}

		stack = append(stack, rect{accumulatedWidth + 1, height})

	}

	return max
}
