# -*- coding: utf-8 -*-
class Solution:
    # @return a ListNode
    def removeNthFromEnd(self, head, n):
        nodes = []
        tmp = head
        while tmp:
            nodes.append(tmp)
            tmp = tmp.next

        rm = nodes[-n]
        if rm == head:
            return head.next
        else:
            pre = nodes[-n-1]
            pre.next = nodes[-n].next
        return head

