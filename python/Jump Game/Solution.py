# -*- coding: utf-8 -*-

class Solution:
    # @param A, a list of integers
    # @return a boolean
    def canJump(self, A):
        lastIndex = len(A) - 1
        i = lastIndex - 1
        while i >= 0:
            if i + A[i] >= lastIndex:
                lastIndex = i
            i -= 1
        return lastIndex == 0