# -*- coding: utf-8 -*-

class Solution:
    # @return a string
    def longestCommonPrefix(self, strs):
        if not strs:
            return ''
        common = strs[0]
        for x in strs:
            tempLength = len(common)
            if len(x) < tempLength:
                tempLength = len(x)
                common,x = x,common
            for i in xrange(1,tempLength+1):
                if common[0:i] != x[0:i]:
                    if i > 1:
                        common = common[0:i-1]
                    else:
                        common = ''
                    break;

            if not common:
                break
        return common