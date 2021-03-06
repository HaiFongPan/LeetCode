package leetcode

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	p0 := nums[0]
	p1 := Max(p0, nums[1])
	for i := 2; i < len(nums); i++ {
		p1, p0 = Max(p1, p0+nums[i]), p1
	}
	return p1
}

func rob2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	temp := 0
	temp1 := nums[1]

	dp := nums[0]
	dp1 := Max(dp, nums[1])

	max := Max(temp1, dp1)
	for i := 2; i < len(nums); i++ {
		if i == len(nums)-1 {
			max = Max(temp+nums[i], dp1)
		}
		dp1, dp = Max(dp1, dp+nums[i]), dp1
		temp1, temp = Max(temp1, temp+nums[i]), temp1
	}

	return max
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
