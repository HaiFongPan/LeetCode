# -*- coding: utf-8 -*-

class Solution:
    # @return a boolean
    leftBrackets = ['{','(','[']
    rightBrackets = ['}',')',']']
    def isValid(self, s):
        leftStack = []
        for bracket in s:
            if bracket in self.leftBrackets:
                leftStack.append(bracket)
            elif bracket in self.rightBrackets:
                index = self.rightBrackets.index(bracket)
                if leftStack and leftStack.pop() == self.leftBrackets[index]:
                    pass
                else:
                    return False
            else:
                return False
        if len(s) == 0 or leftStack:
            return False
        return True