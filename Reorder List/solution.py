# -*- coding: utf-8 -*-

class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution:
    # @param head, a ListNode
    # @return nothing
    def reorderList(self, head):
      if not head:
        return head
      slow = head
      fast = head
      # split into two list by the middle one
      while fast and fast.next:
        slow = slow.next
        fast = fast.next.next

      L1 = head
      L2 = slow.next
      slow.next = None

      if not L2:
        return L1
      if not L2.next:
        L2.next = L1.next
        L1.next = L2
      else:
        pointL1 = L1
        pointL2 = L2
        while pointL2:
          pointL2 = L2
          while pointL2.next and pointL2.next.next:
            pointL2 = pointL2.next

          if pointL2.next:
            tempNode = pointL2.next
            pointL2.next = None
            tempNode.next = pointL1.next
            pointL1.next = tempNode
          elif pointL2:
            tempNode = pointL2
            pointL2 = None
            tempNode.next = pointL1.next
            pointL1.next = tempNode

          pointL1 = pointL1.next.next
      return L1

def printValue(node):
    while node != None:
      print node.val
      node = node.next


# if __name__ == '__main__':
#   a = ListNode(6)
#   b = ListNode(4)
#   c = ListNode(5)
#   d = ListNode(10)
#   e = ListNode(7)
#   f = ListNode(3)
#   g = ListNode(2)
#   a.next = b
#   b.next = c
#   c.next = d
#   # d.next = e
#   # e.next = f
#   # f.next = g
#   s = Solution()
#   x = s.reorderList(a)
#   printValue(x)
