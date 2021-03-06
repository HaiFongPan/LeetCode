package leetcode

import (
	"fmt"
	"sort"
)

type Person struct {
	h int
	w int
}

type Persons []Person

func (p Persons) Len() int {
	return len(p)
}

func (p Persons) Less(i, j int) bool {
	if p[i].h == p[j].h {
		return p[i].w > p[j].w
	}
	return p[i].h < p[j].h
}

func (p Persons) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func bestSeqAtIndex(height []int, weight []int) int {
	if len(height) <= 1 {
		return len(height)
	}
	plist := Persons{}
	for i := 0; i < len(height); i++ {
		plist = append(plist, Person{h: height[i], w: weight[i]})
	}
	sort.Sort(plist)
	dp := Persons{}
	for i := 0; i < len(plist); i++ {
		current := plist[i]
		if i == 0 || (dp[len(dp)-1].w < current.w && dp[len(dp)-1].h < current.h) {
			dp = append(dp, current)
		} else {
			l, r := 0, len(dp)-1
			for l < r {
				mid := (l + r) >> 1
				if dp[mid].w < current.w && dp[mid].h < current.h {
					l = mid + 1
				} else {
					r = mid
				}
			}
			dp[r] = current
		}
	}
	fmt.Printf("len %d", len(dp))
	return len(dp)
}
