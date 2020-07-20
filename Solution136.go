package leetcode

func singleNumber(nums []int) int {
	single := 0
	for i := 0; i < len(nums); i++ {
		single ^= nums[i]
	}

	return single
}

func singleNumber2(nums []int) int {
	bit := make(map[int]int)
	for _, n := range nums {
		bit[n] = bit[n] + 1
	}

	for k, v := range bit {
		if v == 1 {
			return k
		}
	}
	return nums[len(nums)-1]
}
