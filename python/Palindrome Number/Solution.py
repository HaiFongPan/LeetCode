# -*- coding: utf-8 -*-

class Solution:
    # @return a boolean
    def isPalindrome(self, x):
        if x < 0:
            return False;
        length = 0
        tmp = x
        while tmp:
            tmp /= 10
            length += 1
        while length > 1:
            head = x / (10**(length - 1))
            end = x % 10
            if head != end:
                return False
            x %= 10**(length-1)
            x /=10
            length -= 2
        return True