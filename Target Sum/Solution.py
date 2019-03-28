# -*- coding: utf-8 -*-
class Solution(object):
    def findTargetSumWays(self, nums, s):

        sum = 0;
        for num in nums:
            sum = sum + num;

        if s > sum:
            return 0

        coln = sum * 2 + 1;
        rown = len(nums) + 1;

        dp = [[0 for i in range(coln)] for i in range(rown)];
        dp[0][sum] = 1
        target = sum + s;
        for i in xrange(1,rown):
            for j in reversed(xrange(0, coln)):
                if j >= nums[i-1]:
                    dp[i][j] += dp[i-1][j-nums[i-1]]
                if j + nums[i-1] < coln:
                    dp[i][j] += dp[i-1][j+nums[i-1]]
        return dp[rown-1][target]


if __name__ == '__main__':
    s = Solution()
    print s.findTargetSumWays([1,1,1,1,1], 3)
    print s.findTargetSumWays([0,0,0,0,0,0,0,0,1], 1)
    print s.findTargetSumWays([0,0,0,0,0,0,0], 0)

