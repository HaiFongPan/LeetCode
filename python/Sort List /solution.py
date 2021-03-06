# -*- coding: utf-8 -*-

class ListNode:
  def __init__(self, x):
    self.val = x
    self.next = None

class Solution:
    # @param head, a ListNode
    # @return a ListNode
  def sortList(self, head):
    if head == None or head.next == None:
      return head
    slow = head
    fast = head
    while fast.next != None and fast.next.next != None:
      slow = slow.next
      fast = fast.next.next

    L1 = head
    L2 = slow.next
    slow.next = None

    r1 = self.sortList(L1)
    r2 = self.sortList(L2)
    if r1.val <= r2.val:
      head = r1
      r1 = r1.next
    else:
      head = r2
      r2 = r2.next

    headTemp = head

    while r1 != None and r2 != None:
      if r1.val <= r2.val:
        headTemp.next = r1
        r1 = r1.next
      else:
        headTemp.next = r2
        r2 = r2.next
      headTemp = headTemp.next

    if r1 != None:
      headTemp.next = r1
    else:
      headTemp.next = r2

    return head

