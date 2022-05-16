package leetcode

func maximumScore(nums []int, k int) int {
	min := make([]int, len(nums))
	min[k] = nums[k]

	for i := k - 1; i >= 0; i-- {
		min[i] = minf(nums[i], min[i+1])
	}

	for i := k + 1; i < len(nums); i++ {
		min[i] = minf(nums[i], min[i-1])
	}

	l, r, ans := 0, len(nums)-1, 0
	for l <= k && r >= k {
		ans = maxf(ans, (r-l+1)*minf(min[l], min[r]))
		if min[l] < min[r] {
			l = l + 1
		} else {
			r = r - 1
		}
	}

	return ans
}

func minf(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func maxf(x, y int) int {
	if x > y {
		return x
	}
	return y
}
