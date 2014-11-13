
# -*- coding: utf-8 -*-

class Solution:
    # @param s, a string
    # @return a string
    def reverseWords(self, s):
      inputStr = s.split(' ')
      outputStr = ''
      print len(inputStr)
      for x in inputStr[::-1]:
          if x != '':
            outputStr = outputStr + x + ' '
      return outputStr[0:-1]
