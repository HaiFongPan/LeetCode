package pri.leetcode.leetcode.editor.cn;
//给定一个只包含 '(' 和 ')' 的字符串，找出最长的包含有效括号的子串的长度。
//
// 示例 1: 
//
// 输入: "(()"
//输出: 2
//解释: 最长有效括号子串为 "()"
// 
//
// 示例 2: 
//
// 输入: ")()())"
//输出: 4
//解释: 最长有效括号子串为 "()()"
// 
// Related Topics 字符串 动态规划

// ()(()) = 6

import java.util.Stack;

//leetcode submit region begin(Prohibit modification and deletion)
class Solution32 {
    public int longestValidParentheses(String s) {
        Stack<Integer> leftParentheses = new Stack<>();
        if (s.length() < 2) {
            return 0;
        }
        int[] state = new int[s.length()];
        int max = 0;
        for (int i = 0; i < s.length(); i++) {
            if (s.charAt(i) == '(') {
                leftParentheses.push(i);
                continue;
            }

            if (leftParentheses.empty()) {
                continue;
            }

            Integer left = leftParentheses.pop();
            if (i - left > 1) {
                state[i] = (left > 0 ? state[left - 1] : 0) + i - left + 1;
            } else {
                state[i] = (i > 1 ? state[i - 2] : 0) + 2;
            }
            if (state[i] > max) {
                max = state[i];
            }
        }
        return max;
    }

    public static void main(String[] args) {
        Solution32 solution32 = new Solution32();
        System.out.println(solution32.longestValidParentheses(")(((((()())()()))()(()))("));
        System.out.println(solution32.longestValidParentheses(")()())"));
        System.out.println(solution32.longestValidParentheses("()(())"));
    }
}
//leetcode submit region end(Prohibit modification and deletion)
