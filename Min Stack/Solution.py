# -*- coding: utf-8 -*-

class MinStack:
    def __init__(self):
        self.stack = []
        self.lastMin = []

    # @param x, an integer
    # @return an integer
    def push(self, x):
        if len(self.stack) == 0:
            self.lastMin.append(x)
        elif x <= self.lastMin[-1]:
            self.lastMin.append(x)
            # pass
        self.stack.append(x)
        return x

    # @return nothing
    def pop(self):
        if len(self.stack) > 0:
            if self.stack.pop() == self.lastMin[-1]:
                self.lastMin.pop()

    # @return an integer
    def top(self):
        if len(self.stack) == 0:
            return None
        return self.stack[len(self.stack) - 1]

    # @return an integer
    def getMin(self):
        if len(self.lastMin) > 0:
            return self.lastMin[-1]
        return None


