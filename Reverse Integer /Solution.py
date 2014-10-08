# -*- coding: utf-8 -*-

class Solution:
    # @return an integer
    def reverse(self, x):
        reverseX = 0
        needM = False
        if x < 0:
            x = -x
            needM = True
        while x:
            reverseX = reverseX * 10 + x % 10
            x = x/10

        if needM:
            reverseX = -reverseX
        return reverseX