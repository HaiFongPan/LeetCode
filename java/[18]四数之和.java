package pri.leetcode.leetcode.editor.cn;
//给定一个包含 n 个整数的数组 nums 和一个目标值 target，判断 nums 中是否存在四个元素 a，b，c 和 d ，使得 a + b + c + d 的值与 target 相等？找出所有满足条件且不重复的四元组。
//
// 注意： 
//
// 答案中不可以包含重复的四元组。 
//
// 示例： 
//
// 给定数组 nums = [1, 0, -1, 0, -2, 2]，和 target = 0。
//
//满足要求的四元组集合为：
//[
//  [-1,  0, 0, 1],
//  [-2, -1, 1, 2],
//  [-2,  0, 0, 2]
//]
// 
// Related Topics 数组 哈希表 双指针

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

//leetcode submit region begin(Prohibit modification and deletion)
class Solution18 {
    public List<List<Integer>> fourSum(int[] nums, int target) {
        List<List<Integer>> result = new ArrayList<>();
        if (nums.length < 4) {
            return result;
        }

        Arrays.sort(nums);
        for (int i = 0; i < nums.length - 3; i++) {
            if (i == 0 || nums[i] != nums[i - 1]) {
                for (int j = i + 1; j < nums.length - 2; j++) {
                    if (j == i + 1 || nums[j] != nums[j - 1]) {
                        int prev = j + 1, tail = nums.length - 1, sum = target - nums[i] - nums[j];
                        while (prev < tail) {
                            if (nums[prev] + nums[tail] == sum) {
                                result.add(Arrays.asList(nums[i], nums[j], nums[prev], nums[tail]));
                                while (prev < tail && nums[prev] == nums[prev + 1]) prev++;
                                while (prev < tail && nums[tail] == nums[tail - 1]) tail--;
                                prev++;
                                tail--;
                            } else if (nums[prev] + nums[tail] < sum) {
                                prev++;
                            } else {
                                tail--;
                            }
                        }
                    }
                }
            }
        }

        return result;
    }

    public static void main(String[] args) {
        Solution18 solution18 = new Solution18();
//[-1,-5,-5,-3,2,5,0,4]
        //-7
        System.out.println(solution18.fourSum(new int[]{-1,-5,-5,-3,2,5,0,4}, -7));
    }
}
//leetcode submit region end(Prohibit modification and deletion)
