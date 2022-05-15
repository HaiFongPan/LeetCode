package leetcode

func maxSubArray(nums []int) int {
	// sum[i] - sum[j] = [j+1 ... i]
	sum := make([]int, len(nums))
	sum[0] = nums[0]
	min, max := 0, sum[0]
	if sum[0] < 0 {
		min = sum[0]
	}

	for i := 1; i < len(nums); i++ {
		sum[i] = nums[i] + sum[i-1]
		if sum[i]-min > max {
			max = sum[i] - min
		}

		if sum[i] < min {
			min = sum[i]
		}
	}

	return max
}
