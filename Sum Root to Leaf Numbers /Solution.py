# -*- coding: utf-8 -*-

# Definition for a  binary tree node
class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Solution:
    # @param root, a tree node
    # @return an integer
    def sumNumbers(self, root):
      result = []
      sumr = 0
      self.func(root, '', result)
      for x in result:
        sumr += int(x)
      return sumr

    def func(self, root, leafNum, result):
      if root:
        rootNum = leafNum + str(root.val)
        if root.left:
          self.func(root.left, rootNum, result)

        if root.right:
          self.func(root.right, rootNum, result)
        if not root.left and not root.right:
          result.append(rootNum)

# if __name__ == '__main__':
#   leafA = TreeNode(1)
#   leafB = TreeNode(4)
#   leafC = TreeNode(5)
#   leafD = TreeNode(7)
#   leafE = TreeNode(8)
#   rootA = TreeNode(2)
#   rootB = TreeNode(3)

#   rootA.left = leafA
#   rootA.right = rootB
#   rootB.left = leafB
#   rootB.right = leafC
#   leafA.left = leafD
#   leafA.right = leafE

#   s = Solution()
#   print s.sumNumbers(rootA)