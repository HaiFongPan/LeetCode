package leetcode

import "fmt"

type TrieN745 struct {
	Children []*TrieN745
	Weight   int
}

func (this *TrieN745) insert(word string, weight int) {
	for j := 0; j <= len(word); j++ {
		mix := word[j:] + "{" + word
		fmt.Printf("process word %s\n", mix)
		point := this
		for i := 0; i < len(mix); i++ {
			index := mix[i] - 'a'
			if point.Children[index] == nil {
				node := &TrieN745{Weight: weight, Children: make([]*TrieN745, 27)}
				point.Children[index] = node
			}
			point.Children[index].Weight = weight
			point = point.Children[index]
		}
	}
}

func (this *TrieN745) search(prefix string) int {
	fmt.Printf("searching %s\n", prefix)
	point := this
	for i := 0; i < len(prefix); i++ {
		index := prefix[i] - 'a'
		fmt.Printf("searching %d\n", index)
		if point.Children[index] == nil {
			return -1
		}
		point = point.Children[index]
	}
	return point.Weight
}

type WordFilter struct {
	Root *TrieN745
}

func Constructor(words []string) WordFilter {
	root := &TrieN745{Weight: -1, Children: make([]*TrieN745, 27)}
	for k, v := range words {
		root.insert(v, k)
	}
	filter := WordFilter{Root: root}
	return filter
}

func (this *WordFilter) F(prefix string, suffix string) int {
	return this.Root.search(suffix + "{" + prefix)
}

/**
 * Your WordFilter object will be instantiated and called as such:
 * obj := Constructor(words);
 * param_1 := obj.F(prefix,suffix);
 */
