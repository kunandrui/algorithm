package main

import "fmt"

func main() {
	s := "aba"
	fmt.Println(longestPalindrome(s))
}

// 回文串
func isPalindrome(s string) bool {
	n := len(s)

	var ns []byte

	for i := 0; i < n; i++ {
		if s[i] >= 'a' && s[i] <= 'z' || s[i] >= '0' && s[i] <= '9' {
			ns = append(ns, s[i])
		} else if s[i] >= 'A' && s[i] <= 'Z' {
			ns = append(ns, 'a'+(s[i]-'A'))
		} else {
			continue
		}
	}
	fmt.Println(string(ns))
	l := 0
	r := len(ns) - 1
	for l < r {
		if ns[l] == ns[r] {
			l++
			r--
		} else {
			return false
		}
	}

	return true
}

// 给定一个非空字符串 s，最多删除一个字符。判断是否能成为回文字符串。
func validPalindrome(s string) bool {
	l := 0
	r := len(s) - 1

	return check(s, l, r, 1)
}

func check(s string, l, r, times int) bool {
	for l < r {
		if s[l] != s[r] {
			return times > 0 && check(s, l+1, r, times-1) || check(s, l, r-1, times-1)
		}
		l++
		r--
	}
	return true
}
func longestPalindrome(s string) string {
	n := len(s)
	ansLen := 0
	start := 0
	for i := 0; i < n; i++ {
		l := i - 1
		r := i + 1
		for l >= 0 && r < n && s[l] == s[r] {
			l--
			r++
		}
		if r-l-1 > ansLen {
			ansLen = r - l - 1
			start = l + 1
		}
	}

	for i := 0; i < n; i++ {
		l := i
		r := i + 1
		for l >= 0 && r < n && s[l] == s[r] {
			l--
			r++
		}
		if r-l-1 > ansLen {
			ansLen = r - l - 1
			start = l + 1
		}
	}
	fmt.Println(ansLen)
	return s[start : start+ansLen]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
