# -*- coding: utf-8 -*-
# Definition for singly-linked list.
class ListNode:
  def __init__(self, x):
    self.val = x
    self.next = None

class Solution:
  # @param head, a ListNode
  # @return a ListNode
  def insertionSortList(self, head):
    if head == None:
      return head

    result = ListNode(head.val)
    result.next = None
    nextNode = head.next

    while nextNode != None:
      pre = None
      now = result
      while now != None and nextNode.val >= now.val:
        pre = now
        now = now.next

      newNode = ListNode(nextNode.val)
      if now == None:
        pre.next = newNode
      elif pre == None:
        newNode.next = now
        result = newNode
      else:
        pre.next = newNode
        newNode.next = now
      nextNode = nextNode.next

    return result