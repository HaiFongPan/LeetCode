# -*- coding: utf-8 -*-

class Solution:
    # @return a list of integers
    def getRow(self, rowIndex):
      if rowIndex == 0:
        return [1]
      if rowIndex == 1:
        return [1,1]
      arr = [1,1]
      for j in xrange(2, rowIndex+1):
        tmp = arr[0]
        for i in xrange(1, len(arr)):
          tmp1 = arr[i]
          arr[i] = arr[i] + tmp
          tmp = tmp1
        arr.append(1)
      return arr
