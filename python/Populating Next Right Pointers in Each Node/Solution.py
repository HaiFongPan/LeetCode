# -*- coding: utf-8 -*-
# same anwser :Populating Next Right Pointers in Each Node II 
class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None
        self.next = None

class Solution:
    # @param root, a tree node
    # @return nothing
    def connect(self, root):
        if root:
            currentLevelNode = [root]
            self.func(currentLevelNode)
        return root

    def func(self, currentLevelNode):
        while len(currentLevelNode) > 0:
            currentDepthSize = len(currentLevelNode)
            for i in xrange(0,currentDepthSize - 1):
                currentLevelNode[i].next = currentLevelNode[i + 1]
            currentLevelNode[currentDepthSize - 1].next = None
            for i in xrange(0,currentDepthSize):
                if currentLevelNode[i].left:
                    currentLevelNode.append(currentLevelNode[i].left)
                if currentLevelNode[i].right:
                    currentLevelNode.append(currentLevelNode[i].right)
            currentLevelNode = currentLevelNode[currentDepthSize:]