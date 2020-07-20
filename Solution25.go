// Package leetcode provides ...
package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	length := 0
	temp := head
	for temp != nil {
		length++
		temp = temp.Next
	}

	return reverseKGroup2(head, k, length)
}

func reverseKGroup2(head *ListNode, k int, l int) *ListNode {
	if l < k {
		return head
	}

	var temp *ListNode
	prev := head
	cur := head.Next

	for i := 1; cur != nil && i < k; i++ {
		temp = cur.Next
		cur.Next = prev
		prev = cur
		cur = temp
	}
	head.Next = reverseKGroup2(cur, k, l-k)
	return prev

}
