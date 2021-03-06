# -*- coding: utf-8 -*-

class Solution:
    # @return a string
    def countAndSay(self, n):

        currentCAS = '1'
        for i in xrange(1,n):
            currentCAS = self.getSayFromLast(currentCAS)
        return currentCAS

    def getSayFromLast(self, countandsay):
        last = countandsay[0]
        count = 0
        value = ''
        for cha in countandsay:
            if cha == last:
                count += 1
            else:
                value += str(count)+last
                count = 1
            last = cha
        value += str(count)+countandsay[-1]
        return value
