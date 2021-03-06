package pri.leetcode.leetcode.editor.cn;
//假设按照升序排序的数组在预先未知的某个点上进行了旋转。
//
// ( 例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] )。 
//
// 搜索一个给定的目标值，如果数组中存在这个目标值，则返回它的索引，否则返回 -1 。 
//
// 你可以假设数组中不存在重复的元素。 
//
// 你的算法时间复杂度必须是 O(log n) 级别。 
//
// 示例 1: 
//
// 输入: nums = [4,5,6,7,0,1,2], target = 0
//输出: 4
// 
//
// 示例 2: 
//
// 输入: nums = [4,5,6,7,0,1,2], target = 3
//输出: -1 
// Related Topics 数组 二分查找

//leetcode submit region begin(Prohibit modification and deletion)
class Solution33 {
    public int search(int[] nums, int target) {
        int length = nums.length;
        if (length == 0) {
            return -1;
        }
        if (length == 1) {
            return nums[0] == target ? 0 : -1;
        }

        int rotate = 0;
        for (int i = 1; i < length; i++) {
            if (nums[i] < nums[i - 1]) {
                rotate = i;
                break;
            }
        }

        int low = 0, high = length - 1;
        while (low <= high) {
            int mid = low + ((high - low) >> 1);
            int realIndex = realIndex(rotate, mid, length);
            if (nums[realIndex] == target) {
                return realIndex;
            }
            if (nums[realIndex] > target) {
                high = mid - 1;
            } else {
                low = mid + 1;
            }
        }

        return -1;
    }

    private int realIndex(int rotate, int mid, int length) {
        return (rotate + mid) % length;
    }

    public static void main(String[] args) {
        Solution33 s = new Solution33();
        int[] nums = new int[] {4,5,6,7,0,1,2};
        System.out.println(s.search(nums, 0));
    }
}

//leetcode submit region end(Prohibit modification and deletion)
