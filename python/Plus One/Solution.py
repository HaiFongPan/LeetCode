# -*- coding: utf-8 -*-

class Solution(object):
    """docstring for Solution"""
    def plusOne(self, digits):
        digits[-1] += 1
        for i in xrange(1,len(digits) + 1):
            if digits[-i] == 10 and i < len(digits):
                digits[-i] = 0
                digits[-i-1] += 1
        if digits[0] == 10:
            digits[0] = 0
            digits.insert(0,1)

        return digits