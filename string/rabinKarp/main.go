package rabinKarp

func strStr(haystack string, needle string) int {
	n := len(haystack)
	m := len(needle)

	H := make([]int, n+1)
	b := 131
	powBM := 1
	// p := int(1e9) + 7
	for i := 1; i <= n; i++ {
		H[i] = H[i-1]*b + int(haystack[i-1]-'a') + 1
	}

	Hneedle := 0
	for i := 0; i < m; i++ {
		Hneedle = Hneedle*b + int(needle[i]-'a') + 1
		powBM *= b
	}

	//s[l...r]的hash值
	for l := 1; l+m-1 <= n; l++ {
		r := l + m - 1
		//H[r]表示前r - 1
		//H[l]表示前l - 1
		hash := H[r] - H[l-1]*powBM
		if hash == Hneedle {
			return l - 1
		}
	}
	return -1
}
