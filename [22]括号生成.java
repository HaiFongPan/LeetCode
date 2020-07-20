package pri.leetcode.leetcode.editor.cn;
//给出 n 代表生成括号的对数，请你写出一个函数，使其能够生成所有可能的并且有效的括号组合。
//
// 例如，给出 n = 3，生成结果为： 
//
// [
//  "((()))",
//  "(()())",
//  "(())()",
//  "()(())",
//  "()()()"
//]
// 
// Related Topics 字符串 回溯算法

import java.util.ArrayList;
import java.util.List;

//leetcode submit region begin(Prohibit modification and deletion)
class Solution22 {
    public List<String> generateParenthesis(int n) {
        List<String> result = new ArrayList<>();
        if (n > 0) {
            generate("(", n - 1, -1, n, result);
        }
        return result;
    }

    private void generate(String p, int leftP, int balance, int n, List<String> result) {
        if (p.length() == n * 2) {
            result.add(p);
            return;
        }

        if (leftP == 0 || balance < 0) {
            generate(p + ')', leftP, balance + 1, n, result);
        }
        if (leftP > 0) {
            generate(p + '(', leftP - 1, balance - 1, n, result);
        }
    }

    public static void main(String[] args) {
        System.out.println(new Solution22().generateParenthesis(3));
    }
}
//leetcode submit region end(Prohibit modification and deletion)
