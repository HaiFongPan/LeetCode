package pri.leetcode.leetcode.editor.cn;
//给定一个字符串 s，将 s 分割成一些子串，使每个子串都是回文串。
//
// 返回符合要求的最少分割次数。 
//
// 示例: 
//
// 输入: "aab"
//输出: 1
//解释: 进行一次分割就可将 s 分割成 ["aa","b"] 这样两个回文子串。
// 
// Related Topics 动态规划

import java.util.ArrayList;
import java.util.Collections;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

//leetcode submit region begin(Prohibit modification and deletion)
class Solution132 {
    public int minCut(String s) {
        if (s.length() <= 1) {
            return 0;
        }

        //Map<Integer, List<Integer>> charIndex = new HashMap<>();
        //for (int i = 0; i < s.length(); i++) {
        //    int index = s.charAt(i) - 'a';
        //    if (charIndex.containsKey(index)) {
        //        charIndex.get(index).add(i);
        //    } else {
        //        List<Integer> list = new ArrayList<>();
        //        list.add(i);
        //        charIndex.put(index, list);
        //    }
        //}

        int[] state = new int[s.length()];
        state[0] = 0;

        for (int j = 1; j < s.length(); j++) {
            state[j] = state[j - 1] + 1;
            for (int i = 0; i < j; i++) {
                if (isPalindrome(s, i, j)) {
                    if (i == 0) {
                        state[j] = 0;
                        break;
                    }

                    state[j] = Math.min(state[i - 1] + 1, state[j]);
                }
            }
        }

        return state[s.length() - 1];
    }

    public boolean isPalindrome(String s, int i, int j) {
        while (i < j) {
            if (s.charAt(i) == s.charAt(j)) {
                i++;
                j--;
                continue;
            }

            return false;
        }
        return true;
    }

    public static void main(String[] args) {
        Solution132 solution132 = new Solution132();
        System.out.println(solution132.minCut("aabbaaaa"));
        System.out.println(solution132.minCut("bab"));
        System.out.println(solution132.minCut("baba"));
        System.out.println(solution132.minCut("cbbbcc"));
        System.out.println(solution132.minCut("cdd"));
    }
}
//leetcode submit region end(Prohibit modification and deletion)
