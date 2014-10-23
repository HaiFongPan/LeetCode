# -*- coding: utf-8 -*-

class Solution:
    # @param s, a string
    # @return a boolean
    def isPalindrome(self, s):
        s = s.lower()
        onlyAlpha = ''
        for char in s:
            if char.isalnum():
                onlyAlpha += char
        re = onlyAlpha[::-1]
        if onlyAlpha == re:
            return True
        return False