# -*- coding: utf-8 -*-

class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Solution:

    def inorderTraversal(self, root):
      result = []
      return self.func(root, result)

    def func(self, root, result):
        if root:
            if root.left != None:
                self.func(root.left, result)

            result.append(root.val)

            if root.right != None:
                self.func(root.right, result)

        return result