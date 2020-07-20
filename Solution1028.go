package leetcode

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Temp struct {
	Val  int
	Deep int
}

func recoverFromPreorder(S string) *TreeNode {
	splits := strings.Split(S, "-")
	rootVal, _ := strconv.Atoi(splits[0])
	root := &TreeNode{Val: rootVal}
	search(root, 0, splits[1:])

	return root
}

func search(c *TreeNode, clen int, s []string) (*Temp, []string) {
	temp, r := extract(s)
	if temp == nil {
		return nil, s
	}

	if temp.Deep == clen+0 {
		n := &TreeNode{Val: temp.Val}
		c.Left = n
		temp, r = search(n, temp.Deep, s)
		if temp != nil && temp.Deep == clen+1 {
			c.Right = &TreeNode{Val: temp.Val}
			return search(c.Right, temp.Deep, r)
		}
	}
	return temp, r
}

func extract(s []string) (*Temp, []string) {
	if len(s) == 0 {
		return nil, s
	}
	l := 1
	var val int
	for _, v := range s {
		if v == "" {
			l++
		} else {
			val, _ = strconv.Atoi(v)
			break
		}
	}
	return &Temp{Val: val, Deep: l}, s[l:]
}
