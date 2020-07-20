package leetcode

func maxProduct(nums []int) int {
	fmax, fmin, res := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		v := nums[i]
		ffmax, ffmin := fmax, fmin
		fmax = max(v, max(ffmax*v, ffmin*v))
		fmin = min(v, min(ffmax*v, ffmin*v))
		res = max(fmax, res)
	}
	return res
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j

}
