# -*- coding: utf-8 -*-

class Solution:
    # @param tokens, a list of string
    # @return an integer
    def evalRPN(self, tokens):
      #operator = ['+', '-', '*', '/']
      stack = []
      for x in tokens:
        if x == '+':
          a = stack.pop()
          b = stack.pop()
          stack.append(int(a)+int(b))
        elif x == '/':
          a = int(stack.pop())
          b = int(stack.pop())
          stack.append(float(b)/a)
        elif x == '*':
          a = stack.pop()
          b = stack.pop()
          stack.append(int(b)*int(a))
        elif x == '-':
          a = stack.pop()
          b = stack.pop()
          stack.append(int(b)-int(a))
        else:
          stack.append(x)
        # print stack
      return int(stack.pop())