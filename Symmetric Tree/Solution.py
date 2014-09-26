# -*- coding: utf-8 -*-

class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Solution:
    # @param root, a tree node
    # @return a boolean
    # 写得比较乱，至少没LTE
    def isSymmetric(self, root):
        if not root:
            return True
        elif not root.left and root.right:
            return False
        elif root.left and not root.right:
            return False
        elif not root.right and not root.left:
            return True
        pNodes = [root.left,root.right]
        cNodes1 = []
        cNodes2 = []
        flag = True
        while len(pNodes) > 0:
            for i in xrange(0,len(pNodes)/2):
                if pNodes[i].val != pNodes[-1-i].val:
                    return False

                if pNodes[i].left and pNodes[-1-i].right:
                    cNodes1.append(pNodes[i].left)
                    cNodes2.append(pNodes[-1-i].right)
                elif not pNodes[i].left and not pNodes[-1-i].right:
                    pass
                else:
                    return False

                if pNodes[i].right and pNodes[-1-i].left:
                    cNodes1.append(pNodes[i].right)
                    cNodes2.append(pNodes[-1-i].left)
                elif not pNodes[i].right and not pNodes[-1-i].left:
                    pass
                else:
                    return False
            cNodes2.reverse()
            cNodes1 += cNodes2
            pNodes = cNodes1
            cNodes2 =[]
            cNodes1 =[]


        return flag


