package leetcode

import "testing"

type pro0303 struct {
	para0303
	ans0303
}

type para0303 struct {
	p1 []int
	i  int
	j  int
}

type ans0303 struct {
	a1 int
}

func Test_(t *testing.T) {
	ques := []pro0303{
		{
			para0303{[]int{1, 2, 3, 4, 5, 6}, 1, 2},
			ans0303{5},
		},
		{
			para0303{[]int{1, 2, 3, 4, 5, 6}, 0, 5},
			ans0303{21},
		},
	}

	for _, q := range ques {
		p := q.para0303

		obj := Constructor(p.p1)
		v := obj.SumRange(p.i, p.j)
		if v != q.ans0303.a1 {
			t.Fatalf("[input]:%v, [output]: %v, [expect]: %v\n", p, v, q.ans0303.a1)
		}
	}
}
