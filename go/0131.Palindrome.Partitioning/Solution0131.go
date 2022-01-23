package leetcode

func partition(s string) [][]string {
	lens := len(s)
	f := make([][]bool, lens)
	ans := [][]string{}
	splits := []string{}

	for i := 0; i < lens; i++ {
		f[i] = make([]bool, lens)
		for j := 0; j < lens; j++ {
			f[i][j] = true
		}
	}

	for i := lens - 1; i >= 0; i-- {
		for j := i + 1; j < lens; j++ {
			f[i][j] = s[i] == s[j] && f[i+1][j-1]
		}
	}
	var dfs func(s string, splitIndex int)
	dfs = func(s string, splitIndex int) {
		if splitIndex == lens {
			ans = append(ans, append([]string(nil), splits...))
			return
		}

		for j := splitIndex; j < lens; j++ {
			if f[splitIndex][j] {
				splits = append(splits, s[splitIndex:j+1])
				dfs(s, j+1)
				splits = splits[:len(splits)-1]
			}
		}
	}

	dfs(s, 0)
	return ans
}
