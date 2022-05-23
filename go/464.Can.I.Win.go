package leetcode

func canIWin(maxChoosableInteger, desiredTotal int) bool {

	if maxChoosableInteger*(maxChoosableInteger+1)/2 < desiredTotal {
		return false
	}

	dp := make([]int, 1<<maxChoosableInteger)

	var search func(int, int) int
	for i := 0; i < len(dp); i++ {
		dp[i] = -1
	}

	search = func(used, currSum int) (ret int) {
		if dp[used] != -1 {
			return dp[used]
		}
		defer func() { dp[used] = ret }()
		for i := 0; i < maxChoosableInteger; i++ {
			if used>>i&1 == 0 && (currSum+i+1 >= desiredTotal || search(used|1<<i, currSum+i+1) == 0) {
				return 1
			}
		}
		return
	}

	return search(0, 0) == 1
}
