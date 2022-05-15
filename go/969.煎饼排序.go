package leetcode

import "fmt"

func pancakeSort(arr []int) (ans []int) {
	k := len(arr)
	point := k - 1
	for point != 0 {
		i := 0
		for arr[i] != point+1 {
			i++
		}
		if i == 0 {
			i = point
		}
		ans = append(ans, i+1)
		for left, right := 0, i; left < right; left, right = left+1, right-1 {
			arr[left], arr[right] = arr[right], arr[left]
		}
		if arr[point] == point+1 {
			point--
		}

	}
	fmt.Printf("arr: %v\n", arr)
	return ans
}
