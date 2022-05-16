package leetcode

func countKDifference(nums []int, k int) (ans int) {
	cnt := map[int]int{}
	for _, num := range nums {
		ans += cnt[num-k] + cnt[num+k]
		cnt[num]++
	}
	return ans
}

//func countKDifference(nums []int, k int) (ans int) {
//	for i := 0; i < len(nums)-1; i++ {
//		for j := i + 1; j < len(nums); j++ {
//			if abs2006(nums[i], nums[j]) == k {
//				ans++
//			}
//		}
//	}
//	return ans
//}

func abs2006(x, y int) int {
	z := x - y
	if z < 0 {
		return -z
	}
	return z
}
