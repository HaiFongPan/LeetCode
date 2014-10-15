# -*- coding: utf-8 -*-
class Solution:
    # @param root, a tree node
    # @param sum, an integer
    # @return a boolean
    def hasPathSum(self, root, sum):
        if not root:
            return False
        return self.func(root, 0, sum, False)

    def func(self, root, sss, sum, found):
        sss += root.val
        if (not root.left and not root.right) or found:
            return sss == sum
        if root.left and not found:
            found = self.func(root.left, sss, sum, found)

        if root.right and not found:
            found = self.func(root.right, sss, sum, found)
        return found
