# -*- coding: utf-8 -*-

class Solution:
    # @param A, a list of integers
    # @return an integer
    def jump(self, A):
        count = 0
        lastIndex = len(A) - 1
        flag = True
        while flag and lastIndex > 0:
            for j in xrange(0,lastIndex):
                if j + A[j] >= lastIndex:
                    lastIndex = j
                    count += 1
                    flag = True
                    break
                flag = False
        return count