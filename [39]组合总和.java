package pri.leetcode.leetcode.editor.cn;
//给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
//
// candidates 中的数字可以无限制重复被选取。 
//
// 说明： 
//
// 
// 所有数字（包括 target）都是正整数。 
// 解集不能包含重复的组合。 
// 
//
// 示例 1: 
//
// 输入: candidates = [2,3,6,7], target = 7,
//所求解集为:
//[
//  [7],
//  [2,2,3]
//]
// 
//
// 示例 2: 
//
// 输入: candidates = [2,3,5], target = 8,
//所求解集为:
//[
//  [2,2,2,2],
//  [2,3,3],
//  [3,5]
//] 
// Related Topics 数组 回溯算法

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

//leetcode submit region begin(Prohibit modification and deletion)
class Solution39 {
    public List<List<Integer>> combinationSum(int[] candidates, int target) {
        Arrays.sort(candidates);
        return func(target, candidates, 0);
    }

    private List<List<Integer>> func(int target, int[] cadidates, int start) {
        List<List<Integer>> funci = new ArrayList<>();
        for (int i = start; i < cadidates.length; i++) {
            int c = cadidates[i];
            int nextTarget = target - c;
            if (nextTarget == 0) {
                // got it
                List<Integer> temp = new ArrayList<>();
                temp.add(c);
                funci.add(temp);
            } else if (nextTarget > 0) {
                List<List<Integer>> fi = func(nextTarget, cadidates, i); // + c;
                if (fi.size() > 0) {
                    for (List<Integer> f : fi) {
                        if (f.size() > 0) {
                            f.add(c);
                            funci.add(f);
                        }
                    }
                }
            } else {
                break;
            }
        }
        return funci;
    }

    public static void main(String[] args) {
        Solution39 solution39 = new Solution39();
        System.out.println(solution39.combinationSum(new int[] {2, 3, 5}, 8));
    }
}
//leetcode submit region end(Prohibit modification and deletion)
