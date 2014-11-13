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