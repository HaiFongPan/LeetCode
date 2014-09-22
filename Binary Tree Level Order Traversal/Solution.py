# -*- coding: utf-8 -*-

class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Solution:
    # @param root, a tree node
    # @return a list of lists of integers
    def levelOrder(self, root):
        if not root:
            return []
        return self.BFS(root)

    def BFS(self, root):
        rootQueue = [root]
        childQueue = []
        result =[]
        currentLevelVal = []
        while len(rootQueue) > 0:
            currentRoot = rootQueue[0]
            currentLevelVal.append(currentRoot.val)
            if currentRoot.left:
                childQueue.append(currentRoot.left)

            if currentRoot.right:
                childQueue.append(currentRoot.right)

            rootQueue = rootQueue[1:]
            if len(rootQueue) == 0:
                result.append(currentLevelVal)
                rootQueue = childQueue
                childQueue = []
                currentLevelVal = []
        # Binary Tree Level Order Traversal II
        # result.reverse()
        return result
