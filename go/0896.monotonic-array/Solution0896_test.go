package leetcode

import (
	"testing"
)

type pro896 struct {
	para896
	ans896
}

type para896 struct {
	p1 []int
}

type ans896 struct {
	a1 bool
}

func Test_0896(t *testing.T) {
	ques := []pro896{
		{
			para896{[]int{1, 2, 3, 4, 5}},
			ans896{true},
		}, {
			para896{[]int{1, 2, 6, 4, 5}},
			ans896{false},
		}, {
			para896{[]int{1, 1, 1, 1, 1}},
			ans896{true},
		}, {
			para896{[]int{5, 3, 2, 4, 1}},
			ans896{false},
		},
	}

	for _, q := range ques {
		p := q.para896
		v := isMonotonic(p.p1)
		if v != q.ans896.a1 {
			t.Fatalf("[input]:%v, [output]: %v, [expect]: %v\n", p, v, q.ans896.a1)
		}
	}
}
