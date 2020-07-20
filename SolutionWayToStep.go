package leetcode

func waysToStep(n int) int {
	dp := []int{1, 2, 4}
	if n > 3 {
		for i := 4; i <= n; i++ {
			dp = append(dp, (dp[i-3]+dp[i-2]+dp[i-1])%100000007)
		}
	}

	return dp[n]
}
