package leetcode

func sortColors(nums []int) {
	a, b, c := 0, 0, 0
	for _, v := range nums {
		if v == 0 {
			a = a + 1
		} else if v == 1 {
			b = b + 1
		} else {
			c = c + 1
		}
	}

	for i := 0; i < len(nums); i++ {
		if i < a {
			nums[i] = 0
		} else if i < a+b {
			nums[i] = 1
		} else {
			nums[i] = 2
		}
	}
}
