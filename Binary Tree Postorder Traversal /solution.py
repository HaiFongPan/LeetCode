# -*- coding: utf-8 -*-
# 后续遍历二叉树 post order


class TreeNode:
  def __init__(self, x):
      self.val = x
      self.left = None
      self.right = None

class Solution:

  def postorderTraversal(self, root):
    result = []
    return self.func(root, result)

  def func(self, root, result):
    if root:
      if root.left != None:
        self.func(root.left, result)
      if root.right != None:
        self.func(root.right, result)

      # append integer
      result.append(root.val)
    return result

# if __name__ == '__main__':
#   leafA = TreeNode(1)
#   leafB = TreeNode(4)
#   leafC = TreeNode(5)
#   rootA = TreeNode(2)
#   rootB = TreeNode(3)

#   rootA.left = leafA
#   rootA.right = rootB
#   rootB.left = leafB
#   rootB.right = leafC

#   s = Solution()
#   print s.postorderTraversal(None)