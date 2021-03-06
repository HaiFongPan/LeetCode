package pri.leetcode.leetcode.editor.cn;
//给定一个无序的整数数组，找到其中最长上升子序列的长度。
//
// 示例: 
//
// 输入: [10,9,2,5,3,7,101,18]
//输出: 4 
//解释: 最长的上升子序列是 [2,3,7,101]，它的长度是 4。 
//
// 说明: 
//
// 
// 可能会有多种最长上升子序列的组合，你只需要输出对应的长度即可。 
// 你算法的时间复杂度应该为 O(n2) 。 
// 
//
// 进阶: 你能将算法的时间复杂度降低到 O(n log n) 吗? 
// Related Topics 二分查找 动态规划



//leetcode submit region begin(Prohibit modification and deletion)
class Solution300 {
    public int lengthOfLIS(int[] nums) {
        int[] state = new int[nums.length];
        for (int i = 0; i < nums.length; i++) {
            state[i] = 1;
        }

        for (int i = 1; i < nums.length; i++) {
            int max = 0;
            for (int j = i - 1; j >= 0; j--) {
                if (nums[j] < nums[i]) {
                    max = max > state[j]? max:state[j];
                }
            }
            state[i] = max + 1;
        }

        int max = 0;
        for (int i1 : state) {
            max = max > i1 ? max : i1;
        }

        return max;
    }

    public static void main(String[] args) {
        System.out.println(new Solution300().lengthOfLIS(new int[]{1,3,6,7,9,4,10,5,6}));
    }
}
//leetcode submit region end(Prohibit modification and deletion)
