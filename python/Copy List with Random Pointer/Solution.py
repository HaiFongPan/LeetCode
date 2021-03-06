# -*- coding: utf-8 -*-

class RandomListNode:
    def __init__(self, x):
        self.label = x
        self.next = None
        self.random = None

class Solution:
    # @param head, a RandomListNode
    # @return a RandomListNode
    def copyRandomList(self, head):
        if not head:
            return None
        connect = {}
        nextNode = head.next
        deepCopy = RandomListNode(head.label)
        connect[head] = deepCopy
        while nextNode:
            deepCopy.next = RandomListNode(nextNode.label)
            connect[nextNode] = deepCopy.next
            nextNode = nextNode.next
            deepCopy = deepCopy.next

        # print connect
        nextNode = head
        deepCopy = connect[head]
        while nextNode:
            if nextNode.random:
                deepCopy.random = connect[nextNode.random]
            else:
                deepCopy.random = None
            deepCopy = deepCopy.next
            nextNode = nextNode.next

        return connect[head]







