package pri.leetcode.leetcode.editor.cn;
//给定三个字符串 s1, s2, s3, 验证 s3 是否是由 s1 和 s2 交错组成的。
//
// 示例 1: 
//
// 输入: s1 = "aabcc", s2 = "dbbca", s3 = "aadbbcbcac"
//输出: true
// 
//
// 示例 2: 
//
// 输入: s1 = "aabcc", s2 = "dbbca", s3 = "aadbbbaccc"
//输出: false 
// Related Topics 字符串 动态规划

//leetcode submit region begin(Prohibit modification and deletion)
class Solution97 {
    public boolean isInterleave(String s1, String s2, String s3) {
        if (s1.length() == 0) {
            return s2.equals(s3);
        }

        if (s2.length() == 0) {
            return s1.equals(s3);
        }

        if (s1.length() + s2.length() != s3.length()) {
            return false;
        }

        int[][] state = new int[s1.length() + 1][s2.length() + 1];
        boolean match = true;
        for (int i = 0; i < s1.length(); i++) {
            boolean temp = s1.charAt(i) == s3.charAt(i);
            if (match && !temp) {
                match = false;
            }
            state[i + 1][0] = match ? 1 : 0;
        }
        match = true;
        for (int i = 0; i < s2.length(); i++) {
            boolean temp = s2.charAt(i) == s3.charAt(i);
            if (match && !temp) {
                match = false;
            }
            state[0][i + 1] = match ? 1 : 0;
        }
        state[1][1] = (s1.charAt(0) == s3.charAt(0) && s2.charAt(0) == s3.charAt(1)) ||
                (s1.charAt(0) == s3.charAt(1) && s2.charAt(0) == s3.charAt(0)) ? 1 : 0;

        for (int i = 0; i < s1.length(); i++) {
            for (int j = 0; j < s2.length(); j++) {
                if (i == 0 && j == 0) {
                    continue;
                }

                int l2 = j + 1;
                int l1 = i + 1;
                if (state[l1 - 1][l2] == 1) {
                    state[l1][l2] = s1.charAt(i) == s3.charAt(i+j+1) ? 1 : 0;
                } else if (state[l1][l2 - 1] == 1) {
                    state[l1][l2] = s2.charAt(j) == s3.charAt(i+j+1) ? 1 : 0;
                }
            }
        }

        return state[s1.length()][s2.length()] == 1;
    }

    public static void main(String[] args) {
        String s1 = "aabcc", s2 = "dbbca", s3 = "aadbbcbcac";
        Solution97 solution97 = new Solution97();
        System.out.println(solution97.isInterleave(s1, s2, s3));
    }
}
//leetcode submit region end(Prohibit modification and deletion)
