# -*- coding: utf-8 -*-

class Solution(object):
    def findMedianSortedArrays(self, nums1, nums2):
        """
        :type nums1: List[int]
        :type nums2: List[int]
        :rtype: float
        """
        merged_list = []
        i = 0
        j = 0
        nums_long = nums1
        nums_short = nums2
        if len(nums1) == 0:
            merged_list = nums2
        elif len(nums2) == 0:
            merged_list = nums1
        else:
            if nums2[-1] > nums1[-1]:
                nums_long = nums2
                nums_short = nums1
            # print nums_long, nums_short
            for x in xrange(0, len(nums_long)):
                for y in xrange(j, len(nums_short)):
                    if (nums_long[x] < nums_short[y]):
                        merged_list.append(nums_long[x])
                        break;
                    else:
                        merged_list.append(nums_short[y])
                        j = y + 1;
                if j == len(nums_short):
                    merged_list.append(nums_long[x])       

        mid = len(merged_list) / 2
        # print len(merged_list), merged_list, mid
        if len(merged_list) % 2 == 0:
            return (merged_list[mid] + merged_list[mid - 1]) / 2.0
        else:
            return merged_list[mid]
        
if __name__ == '__main__':
    s = Solution()
    nums1 = []
    nums2 = [3, 4, 5]
    print s.findMedianSortedArrays(nums1, nums2)          