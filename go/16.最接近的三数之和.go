package leetcode

/*
 * @lc app=leetcode.cn id=16 lang=golang
 *
 * [16] 最接近的三数之和
 */

// @lc code=start
import (
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	result := nums[0] + nums[1] + nums[2]
	min := Abs(result - target)
	for i := 0; i < len(nums); i++ {
		if i == 0 || nums[i] != nums[i-1] {
			left := i + 1
			right := len(nums) - 1
			sum := target - nums[i]
			for left < right {
				lr := nums[left] + nums[right]
				if Abs(sum-lr) < min {
					result = nums[i] + lr
					min = Abs(sum - lr)
				}
				if lr+nums[i] < target {
					left = left + 1
				} else {
					right = right - 1
				}
			}
		}
	}
	return result
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// @lc code=end
