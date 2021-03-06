# -*- coding: utf-8 -*-

class Solution:
  # @return a list of lists of integers
  def generate(self, numRows):
    result = []
    for i in xrange(1,numRows + 1):
      currentRow = []
      for j in xrange(1, i + 1):
        if j > 1 and j < i:
          currentRow.append(result[i - 2][j - 2] + result[i - 2][j - 1])
        else:
          currentRow.append(1)
      result.append(currentRow)
    return result

if __name__ == '__main__':
  print Solution().generate(11)