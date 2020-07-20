package pri.leetcode.leetcode.editor.cn;
//给定两个大小为 m 和 n 的有序数组 nums1 和 nums2。
//
// 请你找出这两个有序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。 
//
// 你可以假设 nums1 和 nums2 不会同时为空。 
//
// 示例 1: 
//
// nums1 = [1, 3]
//nums2 = [2]
//
//则中位数是 2.0
// 
//
// 示例 2: 
//
// nums1 = [1, 2]
//nums2 = [3, 4]
//
//则中位数是 (2 + 3)/2 = 2.5
// 
// Related Topics 数组 二分查找 分治算法

//leetcode submit region begin(Prohibit modification and deletion)
class Solution4 {
    public double findMedianSortedArrays(int[] nums1, int[] nums2) {
        int mid = (nums1.length + nums2.length) / 2;
        int result = 0, temp = 0;
        int p1 = 0, p2 = 0;
        for (int i = 0; i <= mid; i++) {
            if (p1 == nums1.length) {
                temp = nums2[p2++];
            } else if (p2 == nums2.length) {
                temp = nums1[p1++];
            } else {
                if (nums1[p1] <= nums2[p2]) {
                    temp = nums1[p1++];
                } else {
                    temp = nums2[p2++];
                }
            }
            if (i != mid) {
                result = temp;
            }
        }
        return (nums1.length + nums2.length) % 2 != 0 ? temp : (result + temp) / 2.0;
    }

    public static void main(String[] args) {
        Solution4 solution4 = new Solution4();
        System.out.println(solution4.findMedianSortedArrays(new int[] {1, 2}, new int[] {3}));
        System.out.println(solution4.findMedianSortedArrays(new int[] {1, 2}, new int[] {3, 4}));
        System.out.println(solution4.findMedianSortedArrays(new int[] {1, 2}, new int[] {}));
    }
}
//leetcode submit region end(Prohibit modification and deletion)
