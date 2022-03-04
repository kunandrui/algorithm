package twoSum

func twoSum(numbers []int, target int) []int {
	var ans []int
	n := len(numbers)
	l := 0
	r := n - 1
	for l < r {
		if numbers[l]+numbers[r] < target {
			l++
		} else if numbers[l]+numbers[r] > target {
			r--
		} else {
			ans = append(ans, numbers[l], numbers[r])
			break
		}
	}
	return ans
}
