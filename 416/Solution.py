# -*- coding: utf-8 -*-
class Solution(object):
    def canPartition(self, nums):

        sum = 0
        for num in nums:
            sum = sum + num

        if sum % 2 != 0:
            return False

        eq = sum/2
        dp = [0]*(eq+1)
        dp[0] = 1
        print dp
        for num in nums:
            for j in xrange(1,eq+1):
                _j = eq - j + 1
                if _j >= num:
                    dp[_j] = dp[_j-num]
                if _j == eq and dp[_j] == 1:
                    return True
        return False
        # dp = [[0 for i in range(eq+1)] for i in range(len(nums)+1)]
        # print eq

        # for i in xrange(0, len(nums) + 1):
        #     dp[i][0] = 1
        # for i in xrange(1,len(nums)+1):
        #     for j in xrange(1, eq+1):
        #         dp[i][j] = dp[i-1][j]
        #         if j >= nums[i-1]:
        #             dp[i][j] = dp[i-1][j-nums[i-1]]
        #         if j == eq and dp[i][j] == 1:
        #             return True
        # # for i in xrange(0, len(nums) + 1):
        # #     print dp[i]
        # return False

if __name__ == '__main__':
     s = Solution()
     print s.canPartition([1,5,11,5,2])