# -*- coding: utf-8 -*-
class Solution:
    # @param head, a ListNode
    # @return a ListNode
    def deleteDuplicates(self, head):
        if not head:
            return None
        pre = head
        current = head.next
        while current:
            if current.val == pre.val:
                pre.next = current.next
            else:
                pre = pre.next
            current = current.next

        return head
