package main

import "fmt"

func main() {

	fmt.Println(letterCombinations("234"))

}

func letterCombinations(digits string) []string {
	var ans []string
	hash := make(map[byte]string)
	hash['2'] = "abc"
	hash['3'] = "def"
	hash['4'] = "ghi"
	hash['5'] = "jkl"
	hash['6'] = "mno"
	hash['7'] = "pqrs"
	hash['8'] = "tuv"
	hash['9'] = "wxyz"
	var str []byte
	var dfs func(i int)
	dfs = func(i int) {
		if i == len(digits) {
			ans = append(ans, string(append([]byte(nil), str...)))
			return
		}

		for j := 0; j < len(hash[digits[i]]); j++ {
			str = append(str, hash[digits[i]][j])
			dfs(i + 1)
			str = str[0 : len(str)-1]
		}
	}
	dfs(0)
	return ans
}
