# -*- coding: utf-8 -*-

class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution:
    # @return a ListNode
    def addTwoNumbers(self, l1, l2):
        extra = 0
        rhead = l2
        end = None
        extraNode = None
        # 相同长度部分相加
        while l1 and l2:
            v = l1.val + l2.val + extra
            l2.val = v % 10
            extra = v / 10
            l1 = l1.next
            end = l2
            l2 = l2.next
        # 较长 list 部分与进位的计算 extra = 1,list=[9,9,9,9,9,9,9,9,9]
        if l1:
            extraNode = l1
        elif l2:
            extraNode = l2
        end.next = extraNode
        while extraNode and extra:
            v = extraNode.val + extra
            extraNode.val = v % 10
            extra = v / 10
            end = extraNode
            extraNode = extraNode.next

        if extra > 0 :
            end.next = ListNode(extra)
        return rhead

