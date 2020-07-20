package pri.leetcode.leetcode.editor.cn;
//给定一个字符串 S 和一个字符串 T，计算在 S 的子序列中 T 出现的个数。
//
// 一个字符串的一个子序列是指，通过删除一些（也可以不删除）字符且不干扰剩余字符相对位置所组成的新字符串。（例如，"ACE" 是 "ABCDE" 的一个子序列，而 "AEC" 不是） 
//
// 示例 1: 
//
// 输入: S = "rabbbit", T = "rabbit"
//输出: 3
//解释:
//
//如下图所示, 有 3 种可以从 S 中得到 "rabbit" 的方案。
//(上箭头符号 ^ 表示选取的字母)
//
//rabbbit
//^^^^ ^^
//rabbbit
//^^ ^^^^
//rabbbit
//^^^ ^^^
// 
//
// 示例 2: 
//
// 输入: S = "babgbag", T = "bag"
//输出: 5
//解释:
//
//如下图所示, 有 5 种可以从 S 中得到 "bag" 的方案。 
//(上箭头符号 ^ 表示选取的字母)
//
//babgbag
//^^ ^
//babgbag
//^^    ^
//babgbag
//^    ^^
//babgbag
//  ^  ^^
//babgbag
//    ^^^ 
// Related Topics 字符串 动态规划

//leetcode submit region begin(Prohibit modification and deletion)
class Solution115 {
    public int numDistinct(String s, String t) {
        int slen = s.length();
        int tlen = t.length();
        if (slen == 0 || tlen == 0) return 0;
        int[][] state = new int[slen][tlen];
        state[slen - 1][tlen - 1] = s.charAt(slen - 1) == t.charAt(tlen - 1) ? 1 : 0;
        for (int i = slen - 2; i >= 0; i--) {
            if (s.charAt(i) == t.charAt(t.length() - 1)) {
                state[i][tlen - 1] = state[i + 1][tlen - 1] + 1;
            } else {
                state[i][tlen - 1] = state[i + 1][tlen - 1];
            }
        }

        for (int i = slen - 2; i >= 0; i--) {
            for (int j = tlen - 2; slen - i >= tlen - j && j >= 0; j--) {
                if (s.charAt(i) == t.charAt(j)) {
                    state[i][j] = state[i+1][j] + state[i+1][j+1];
                } else {
                    state[i][j] = state[i+1][j];
                }
            }
        }

        return state[0][0];
    }

    public static void main(String[] args) {
        Solution115 solution115 = new Solution115();
        System.out.println(solution115.numDistinct("rabbbit", "rabbit"));
        System.out.println(solution115.numDistinct("babgbag", "bag"));
        System.out.println(solution115.numDistinct("ag", "bag"));
    }
}
//leetcode submit region end(Prohibit modification and deletion)
