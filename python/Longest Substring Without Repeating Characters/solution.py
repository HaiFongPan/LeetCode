# -*- coding: utf-8 -*-

class Solution(object):
    def lengthOfLongestSubstring(self, s):
        """
        :type s: str
        :rtype: int
        """
        start_index = 0
        index_dict = {}
        output = 0
        output_temp = 0
        for i in xrange(0, len(s)):
            if (s[i] in index_dict.keys() and index_dict[s[i]] >= start_index):
                start_index = index_dict[s[i]] + 1
            output = max(i - start_index + 1, output)    
            index_dict[s[i]] = i
        return output

if __name__ == '__main__':
    s = Solution()
    print s.lengthOfLongestSubstring('pwwkew');
    print s.lengthOfLongestSubstring('bbbbb');
    print s.lengthOfLongestSubstring('klsdjaglkjklsdjg084604uglsdkjgkaldshdasjhgl');