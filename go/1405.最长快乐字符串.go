// Package leet provides ...
package leetcode

import "sort"

func longestDiverseString(a int, b int, c int) string {
	ans := []byte{}

	type Cha1405 struct {
		ch  byte
		cnt int
	}

	chs := []Cha1405{{'a', a}, {'b', b}, {'c', c}}
	hasNext := true
	for hasNext {
		sort.Slice(chs, func(i, j int) bool {
			return chs[i].cnt > chs[j].cnt
		})
		hasNext = false

		for i, v := range chs {
			if v.cnt == 0 {
				break
			}

			l := len(ans)
			if l >= 2 && ans[l-1] == v.ch && ans[l-2] == v.ch {
				// already append twice
				continue
			}
			hasNext = true
			ans = append(ans, v.ch)
			chs[i].cnt--
			break
		}
	}

	return string(ans)
}
