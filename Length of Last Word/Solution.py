# -*- coding: utf-8 -*-

class Solution:
    # @param s, a string
    # @return an integer
    def lengthOfLastWord(self, s):
        if not s:
            return 0
        sList = s.strip().split(" ")
        return len(sList[-1])