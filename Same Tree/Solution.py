# -*- coding: utf-8 -*-
class Solution:
    # @param p, a tree node
    # @param q, a tree node
    # @return a boolean
    def isSameTree(self, p, q):
        if not p and not q:
            return True
        if not p and q:
            return False

        return self.func(p, q)

    def func(self, pNode, qNode):
        if not qNode:
            return False
        if pNode.val != qNode.val:
            return False

        isSameNode = True

        if pNode.left:
            isSameNode = self.func(pNode.left, qNode.left)
        elif qNode.left:
            return False

        if  isSameNode and pNode.right:
            isSameNode = self.func(pNode.right, qNode.right)
        elif isSameNode and qNode.right:
            return False

        return isSameNode