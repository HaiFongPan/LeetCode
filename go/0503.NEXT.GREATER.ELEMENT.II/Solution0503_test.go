package leetcode

import "testing"

type pro0503 struct {
	para0503
	ans0503
}

type para0503 struct {
	p1 []int
}

type ans0503 struct {
	a1 []int
}

func Test_(t *testing.T) {
	ques := []pro0503{
		{
			para0503{[]int{1, 2, 1}},
			ans0503{[]int{2, -1, 2}},
		},
		{
			para0503{[]int{1, 1, 1}},
			ans0503{[]int{-1, -1, -1}},
		},
		{
			para0503{[]int{5, 4, 3, 2, 1}},
			ans0503{[]int{-1, 5, 5, 5, 5}},
		},
		{
			para0503{[]int{5, 4}},
			ans0503{[]int{-1, 5}},
		},
		{
			para0503{[]int{5}},
			ans0503{[]int{-1}},
		},
	}

	for _, q := range ques {
		p := q.para0503
		v := nextGreaterElements(p.p1)

		for i, ele := range v {
			if ele != q.ans0503.a1[i] {
				t.Fatalf("[input]:%v, [output]: %v, [expect]: %v\n", p, v, q.ans0503.a1)
			}
		}
	}
}
