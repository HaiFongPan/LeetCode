# -*- coding: utf-8 -*-

class Solution:
    # @param num, a list of integer
    # @return an integer
    def findMin(self, num):
        if not num or len(num) == 0:
            return None
        left = 0
        right = len(num) - 1
        while left < right:
            mid = (left + right)/2
            if num[mid] >= num[left] and num[mid] >= num[right]:
                left = mid + 1
            elif num[mid] >= num[left] and num[mid] < num[right]:
                return num[left]
            else:
                right = mid
        return num[left]

    def findMin2(self, num):
        if not num or len(num) == 0:
            return None
        min = num[0]
        for x in num:
            if x < min:
                min = x
        return min
