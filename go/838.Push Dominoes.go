package leetcode

func pushDominoes(dominoes string) string {
	left := 0
	flag := false
	ans := make([]byte, len(dominoes))
	for i := 0; i < len(dominoes); i++ {
		ans[i] = '.'
		if dominoes[i] == 'R' {
			if flag {
				for ll := left; ll < i; ll++ {
					ans[ll] = 'R'
				}
			}
			flag = true
			left = i
		}

		if dominoes[i] == 'L' {
			if !flag {
				for ll := left; ll <= i; ll++ {
					ans[ll] = 'L'
				}
			} else {
				for ll, rr := left, i; ll <= rr; ll, rr = ll+1, rr-1 {
					if ll == rr {
						ans[ll] = '.'
					} else {
						ans[ll] = 'R'
						ans[rr] = 'L'
					}

				}
			}
			left = i + 1
			flag = false
		}
	}

	if flag {
		for i := left; i < len(dominoes); i++ {
			ans[i] = 'R'
		}
	}

	return string(ans)
}
