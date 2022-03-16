package isMatch

func isMatch(s string, p string) bool {
	m := len(s)
	n := len(p)
	s = " " + s
	p = " " + p

	f := make([][]bool, m+1)
	for i := 0; i <= m; i++ {
		f[i] = make([]bool, n+1)
	}

	f[0][0] = true

	for j := 2; j <= n; j += 2 {
		if p[j] == '*' {
			f[0][j] = true
		} else {
			break
		}
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j] >= 'a' && p[j] <= 'z' {
				f[i][j] = f[i-1][j-1] && s[i] == p[j]
			} else if p[j] == '.' {
				f[i][j] = f[i-1][j-1]
			} else if p[j] == '*' {
				f[i][j] = f[i][j] || f[i][j-2]
				if p[j-1] == '.' || p[j-1] == s[i] {
					f[i][j] = f[i][j] || f[i-1][j]
				}
			}
		}
	}
	return f[m][n]
}
