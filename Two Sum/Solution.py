# -*- coding: utf-8 -*-

class Solution:
    # @return a tuple, (index1, index2)
    def twoSum(self, num, target):
        dictNum = dict.fromkeys(num)
        for i in xrange(0,len(num)):
            if dictNum.has_key(target - num[i]):
                for v in xrange(i+1,len(num)):
                    if num[v] + num[i] ==target:
                        return (i+1,v+1)