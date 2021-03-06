# -*- coding: utf-8 -*-

class MinStack:
    def __init__(self):
        self.stack = []
        self.lastMin = []

    # @param x, an integer
    # @return an integer
    def push(self, x):
        if len(self.stack) == 0 or x <= self.lastMin[-1]:
            self.lastMin.append(x)
        self.stack.append(x)
        return x

    # @return nothing
    def pop(self):
        if len(self.stack) > 0:
            if self.stack.pop() == self.lastMin[-1]:
                self.lastMin.pop()

    # @return an integer
    def top(self):
        return self.stack[-1]

    # @return an integer
    def getMin(self):
        return self.lastMin[-1]


