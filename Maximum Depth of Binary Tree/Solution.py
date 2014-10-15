# -*- coding: utf-8 -*-

class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Solution:
    # @param root, a tree node
    # @return an integer
    def maxDepth(self, root):
        if not root:
            return 0
        r = {'max':0}
        currentDepth = 0
        self.func(root, r, currentDepth)
        return r['max']

    def func(self, root, r, currentDepth):
        currentDepth += 1
        if root.left:
            self.func(root.left, r, currentDepth)

        if root.right:
            self.func(root.right, r, currentDepth)

        if not root.left and not root.right and currentDepth > r['max']:
            r['max'] = currentDepth