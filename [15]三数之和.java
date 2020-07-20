package pri.leetcode.leetcode.editor.cn;
//给定一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？找出所有满足条件且不重复的三元组。
//
// 注意：答案中不可以包含重复的三元组。 
//
// 例如, 给定数组 nums = [-1, 0, 1, 2, -1, -4]，
//
//满足要求的三元组集合为：
//[
//  [-1, 0, 1],
//  [-1, -1, 2]
//]
// 
// Related Topics 数组 双指针

import java.util.ArrayList;
import java.util.Arrays;
import java.util.HashMap;
import java.util.HashSet;
import java.util.List;
import java.util.Map;
import java.util.Set;
import java.util.Stack;

//leetcode submit region begin(Prohibit modification and deletion)
class Solution15 {

    public List<List<Integer>> threeSum(int[] nums) {
        List<List<Integer>> result = new ArrayList<>();
        if (nums.length < 3) {
            return result;
        }

        Arrays.sort(nums);
        for (int i = 0; i < nums.length - 2; i++) {
            if (i == 0 || nums[i] != nums[i - 1]) {
                int prev = i + 1, tail = nums.length - 1, sum = 0 - nums[i];
                while (prev < tail) {
                    if (nums[prev] + nums[tail] == sum) {
                        result.add(Arrays.asList(nums[i], nums[prev], nums[tail]));
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

        return result;
    }

    public static void main(String[] args) {
        Solution15 s = new Solution15();
        long start = System.currentTimeMillis();
        System.out.println(s.threeSum(
                new int[] {1, -1, -1 ,0}));
        System.out.println(System.currentTimeMillis() - start);
    }
}
//leetcode submit region end(Prohibit modification and deletion)
