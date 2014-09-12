# -*- coding: utf-8 -*-

class Solution:
  # @param A, a list of integer
  # @return an integer
  def singleNumber(self, A):
    if len(A) == 1:
      return A[0]
    A.sort()
    time = 1
    for x in xrange(1,len(A)):
      if A[x] == A[x-1]:
        time = time + 1
      elif time == 1:
        return A[x-1]
      else:
        time = 1
    return A[-1]

if __name__ == '__main__':
  s = Solution()
  print s.singleNumber([17,12,5,-6,12,4,17,-5,2,-3,2,4,5,16,-3,-4,15,15,-4,-5,-6])
        