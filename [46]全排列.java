package pri.leetcode.leetcode.editor.cn;
//给定一个没有重复数字的序列，返回其所有可能的全排列。
//
// 示例: 
//
// 输入: [1,2,3]
//输出:
//[
//  [1,2,3],
//  [1,3,2],
//  [2,1,3],
//  [2,3,1],
//  [3,1,2],
//  [3,2,1]
//] 
// Related Topics 回溯算法

import java.util.ArrayList;
import java.util.List;
import java.util.Stack;

//leetcode submit region begin(Prohibit modification and deletion)
class Solution46 {
    private List<List<Integer>> result = new ArrayList<>();

    public List<List<Integer>> permute(int[] nums) {
        if (nums.length == 0) return new ArrayList<>();

        result.clear();
        Stack<Integer> stack = new Stack<>();
        boolean[] visited = new boolean[nums.length];
        permute(nums, visited, stack);
        return result;
    }

    private void permute(int[] nums, boolean[] visited, Stack<Integer> stack) {
        if (stack.size() == nums.length) {
            result.add(new ArrayList<>(stack));
            return;
        }

        for (int i = 0; i < nums.length; i++) {
            if (!visited[i]) {
                visited[i] = true;
                stack.push(nums[i]);
                permute(nums, visited, stack);
                stack.pop();
                visited[i] = false;
            }
        }
    }

    public static void main(String[] args) {
        Solution46 solution46 = new Solution46();
        System.out.println(solution46.permute(new int[]{1,2,3}));
    }
}
//leetcode submit region end(Prohibit modification and deletion)
