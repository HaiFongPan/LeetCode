package pri.leetcode.leetcode.editor.cn;
//给定两个单词 word1 和 word2，计算出将 word1 转换成 word2 所使用的最少操作数 。
//
// 你可以对一个单词进行如下三种操作： 
//
// 
// 插入一个字符 
// 删除一个字符 
// 替换一个字符 
// 
//
// 示例 1: 
//
// 输入: word1 = "horse", word2 = "ros"
//输出: 3
//解释: 
//horse -> rorse (将 'h' 替换为 'r')
//rorse -> rose (删除 'r')
//rose -> ros (删除 'e')
// 
//
// 示例 2: 
//
// 输入: word1 = "intention", word2 = "execution"
//输出: 5
//解释: 
//intention -> inention (删除 't')
//inention -> enention (将 'i' 替换为 'e')
//enention -> exention (将 'n' 替换为 'x')
//exention -> exection (将 'n' 替换为 'c')
//exection -> execution (插入 'u')
// 
// Related Topics 字符串 动态规划

//leetcode submit region begin(Prohibit modification and deletion)
class Solution72 {
    public int minDistance(String word1, String word2) {
        if (word1.length() == 0) return word2.length();
        if (word2.length() == 0) return word1.length();
        int[][] state = new int[word1.length()][word2.length()];
        state[0][0] = word1.charAt(0) == word2.charAt(0) ? 0 : 1;
        for (int i = 1; i < word1.length(); i++) {
            if (word2.charAt(0) == word1.charAt(i)) {
                state[i][0] = i;
            } else {
                state[i][0] = state[i - 1][0] + 1;
            }
        }

        for (int j = 1; j < word2.length(); j++) {
            if (word1.charAt(0) == word2.charAt(j)) {
                state[0][j] = j;
            } else {
                state[0][j] = state[0][j - 1] + 1;
            }
        }

        for (int i = 1; i < word1.length(); i++) {
            for (int j = 1; j < word2.length(); j++) {
                if (word1.charAt(i) == word2.charAt(j)) {
                    state[i][j] = min(state[i - 1][j] + 1, state[i - 1][j - 1], state[i][j - 1] + 1);
                } else {
                    state[i][j] = min(state[i - 1][j] + 1, state[i - 1][j - 1] + 1, state[i][j - 1] + 1);
                }
            }
        }

        return state[word1.length() - 1][word2.length() - 1];
    }

    private int min(int a, int b, int c) {
        return a < b ? (a > c ? c : a) : (b > c ? c : b);
    }

    public static void main(String[] args) {
        Solution72 solution72 = new Solution72();
        System.out.println(solution72.minDistance("horse", "ros"));
        System.out.println(solution72.minDistance("intention", "execution"));
    }
}
//leetcode submit region end(Prohibit modification and deletion)
