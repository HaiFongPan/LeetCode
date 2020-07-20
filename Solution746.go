package leetcode

func minCostClimbingStairs(cost []int) int {
	if len(cost) == 2 {
		return min746(cost[0], cost[1])
	}

	dp1 := cost[0]
	dp2 := min746(dp1, cost[1])

	for i := 2; i < len(cost); i++ {
		v := cost[i]
		dp2, dp1 = min746(dp1+v, dp2+v), dp2
	}
	return min746(dp1, dp2)
}

func min746(x, y int) int {
	if x < y {
		return x
	}
	return y
}
