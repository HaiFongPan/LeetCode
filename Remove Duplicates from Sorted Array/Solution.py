# -*- coding: utf-8 -*-

class Solution:
    # @param a list of integers
    # @return an integer
    def removeDuplicates(self, A):
        if not A:
            return 0
        i = 0
        j = 0
        count = 1
        while j < len(A):
            if A[j] == A[i]:
                j += 1
            else:
                i += 1
                count += 1
                A[i] = A[j]
                j += 1
        return count