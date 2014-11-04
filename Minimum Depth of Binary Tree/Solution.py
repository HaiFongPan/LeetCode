# -*- coding: utf-8 -*-

class Solution:
    # @param root, a tree node
    # @return an integer
    def minDepth(self, root):
        if not root:
            return 0
        r = {'min':-1}
        currentDepth = 0
        self.func(root, r, currentDepth)
        return r['min']

    def func(self, root, r, currentDepth):
        currentDepth += 1
        if root.left:
            self.func(root.left, r, currentDepth)

        if root.right:
            self.func(root.right, r, currentDepth)

        if not root.left and not root.right and (currentDepth < r['min'] or r['min'] == -1):
            r['min'] = currentDepth