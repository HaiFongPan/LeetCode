package pri.leetcode.leetcode.editor.cn;
//给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。
//
// 你的算法时间复杂度必须是 O(log n) 级别。 
//
// 如果数组中不存在目标值，返回 [-1, -1]。 
//
// 示例 1: 
//
// 输入: nums = [5,7,7,8,8,10], target = 8
//输出: [3,4] 
//
// 示例 2: 
//
// 输入: nums = [5,7,7,8,8,10], target = 6
//输出: [-1,-1] 
// Related Topics 数组 二分查找

//leetcode submit region begin(Prohibit modification and deletion)
class Solution34 {
    public int[] searchRange(int[] nums, int target) {
        int[] result = new int[] {-1, -1};
        int find = find(nums, 0, nums.length - 1, target);
        if (find == -1) return result;
        result[0] = find;
        result[1] = find;
        int left = find, right = find;
        left--;
        while (left >= 0 && nums[left] == target) {
            result[0] = left;
            left--;
        }
        right++;
        while (right < nums.length && nums[right] == target) {
            result[1] = right;
            right++;
        }
        return result;
    }

    private int find(int[] nums, int left, int right, int target) {
        if (left > right) return -1;
        if (left == right) {
            return nums[left] == target ? left : -1;
        }
        int mid = (right - left) / 2 + left;
        if (nums[mid] == target) {
            return mid;
        }
        if (nums[mid] > target) {
            return find(nums, left, mid - 1, target);
        } else {
            return find(nums, mid + 1, right, target);
        }
    }

    public static void main(String[] args) {
        Solution34 solution34 = new Solution34();
        int[] ints = solution34.searchRange(new int[] {5, 7, 7, 8, 8, 10}, 6);
        System.out.println(ints[0] + " " + ints[1]);
    }
}
//leetcode submit region end(Prohibit modification and deletion)
