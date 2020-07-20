package pri.leetcode.leetcode.editor.cn;
//给定一个未排序的整数数组，找出其中没有出现的最小的正整数。
//
// 示例 1: 
//
// 输入: [1,2,0]
//输出: 3
// 
//
// 示例 2: 
//
// 输入: [3,4,-1,1]
//输出: 2
// 
//
// 示例 3: 
//
// 输入: [7,8,9,11,12]
//输出: 1
// 
//
// 说明: 
//
// 你的算法的时间复杂度应为O(n)，并且只能使用常数级别的空间。 
// Related Topics 数组



//leetcode submit region begin(Prohibit modification and deletion)
class Solution41 {
    public int firstMissingPositive(int[] nums) {
        int[] temp = new int[nums.length + 2];

        for (int num: nums) {
            if (num <= 0 || num > nums.length)
                continue;
            temp[num] = 1;
        }

        for (int j = 1; j < temp.length;j++) {
            if (temp[j] == 0)
                return j;
        }
        return 0;
    }

    public static void main(String[] args) {
        Solution41 solution41 = new Solution41();
        System.out.println(solution41.firstMissingPositive(new int[]{1,2,0}));
        System.out.println(solution41.firstMissingPositive(new int[]{3,4,-1,1}));
        System.out.println(solution41.firstMissingPositive(new int[]{7,8,9,11,12}));
        System.out.println(solution41.firstMissingPositive(new int[]{1,2,3,4,5}));
    }
}
//leetcode submit region end(Prohibit modification and deletion)
