package leetcode

func isBipartite(graph [][]int) bool {
	n := len(graph)
	toggle := make([]int, n)
	scanned := make([]int, n)
	for i := 0; i < n; i++ {
		toggle[i] = 0
		scanned[i] = 0
	}

	for i := 0; i < n; i++ {
		if len(graph) == 0 {
			continue
		}
		if toggle[i] != 0 {
			continue
		}
		toggle[i] = 1
		b := t(graph, i, -1, toggle, scanned)
		if !b {
			return false
		}
	}
	return true
}

func t(graph [][]int, row int, toggled int, toggle []int, scanned []int) bool {
	if scanned[row] == 1 {
		return true
	}

	for _, link := range graph[row] {
		if toggle[link] == toggled {
			continue
		}
		if toggle[link] != 0 {
			return false
		}

		toggle[link] = toggled
		b := t(graph, link, toggled*-1, toggle, scanned)
		if !b {
			return false
		}
	}
	scanned[row] = 1
	return true
}
