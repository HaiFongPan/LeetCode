package leetcode

func nextGreaterElements(elements []int) []int {
	lens := len(elements)
	r := make([]int, lens)
	stack := []int{}
	for i := 0; i < lens; i++ {
		r[i] = -1
	}

	for i := 0; i < 2*lens-1; i++ {
		for len(stack) > 0 && elements[stack[len(stack)-1]] < elements[i%lens] {
			r[stack[len(stack)-1]] = elements[i%lens]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i%lens)
	}
	return r
}

//func nextGreaterElements(nums []int) []int {
//    n := len(nums)
//    ans := make([]int, n)
//    for i := range ans {
//        ans[i] = -1
//    }
//    stack := []int{}
//    for i := 0; i < n*2-1; i++ {
//        for len(stack) > 0 && nums[stack[len(stack)-1]] < nums[i%n] {
//            ans[stack[len(stack)-1]] = nums[i%n]
//            stack = stack[:len(stack)-1]
//        }
//        stack = append(stack, i%n)
//    }
//    return ans
//}
