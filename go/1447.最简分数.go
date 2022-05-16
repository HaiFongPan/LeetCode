package leetcode

import "fmt"

func simplifiedFractions(n int) (ans []string) {
	for i := 1; i < n; i++ {
		for j := i + 1; j <= n; j++ {
			if commonDivisor1447(i, j) == 1 {
				ans = append(ans, fmt.Sprintf("%d/%d", i, j))
			}
		}
	}
	return ans
}

func commonDivisor1447(x, y int) int {
	if x, y = y, x%y; y != 0 {
		return commonDivisor1447(x, y)
	}
	return x
}
