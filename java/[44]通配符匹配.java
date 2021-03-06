package pri.leetcode.leetcode.editor.cn;
//给定一个字符串 (s) 和一个字符模式 (p) ，实现一个支持 '?' 和 '*' 的通配符匹配。
//
// '?' 可以匹配任何单个字符。
//'*' 可以匹配任意字符串（包括空字符串）。
// 
//
// 两个字符串完全匹配才算匹配成功。 
//
// 说明: 
//
// 
// s 可能为空，且只包含从 a-z 的小写字母。 
// p 可能为空，且只包含从 a-z 的小写字母，以及字符 ? 和 *。 
// 
//
// 示例 1: 
//
// 输入:
//s = "aa"
//p = "a"
//输出: false
//解释: "a" 无法匹配 "aa" 整个字符串。 
//
// 示例 2: 
//
// 输入:
//s = "aa"
//p = "*"
//输出: true
//解释: '*' 可以匹配任意字符串。
// 
//
// 示例 3: 
//
// 输入:
//s = "cb"
//p = "?a"
//输出: false
//解释: '?' 可以匹配 'c', 但第二个 'a' 无法匹配 'b'。
// 
//
// 示例 4: 
//
// 输入:
//s = "adceb"
//p = "*a*b"
//输出: true
//解释: 第一个 '*' 可以匹配空字符串, 第二个 '*' 可以匹配字符串 "dce".
// 
//
// 示例 5: 
//
// 输入:
//s = "acdcb"
//p = "a*c?b"
//输入: false 
// Related Topics 贪心算法 字符串 动态规划 回溯算法

import java.util.ArrayList;
import java.util.List;

//leetcode submit region begin(Prohibit modification and deletion)
class Solution44 {
    public boolean isMatch(String text, String pattern) {
        if (pattern.length() == 0) {
            return text.length() == 0;
        }
        final String[] split = pattern.split("\\*");
        boolean canNotMatch = pattern.charAt(0) == '*';
        boolean canNotMatchLength = pattern.charAt(pattern.length() - 1) == '*';
        int matchedCount = 0;
        int temp = 0;
        for (int j = 0; j < split.length && temp < text.length(); ) {
            boolean matched = true;
            int i = temp;
            if (j != split.length - 1 || canNotMatchLength) {
                for (int sj = 0; sj < split[j].length() && matched; sj++) {
                    if (text.charAt(i++) != split[j].charAt(sj)
                            && split[j].charAt(sj) != '?') {
                        matched = false;
                    }

                    if (i >= text.length() && sj != split[j].length() - 1) {
                        matched = false;
                    }
                }

            } else {
                i = text.length() - 1;
                for (int sj = split[j].length() - 1; sj >= 0; sj--) {
                    if (text.charAt(i--) != split[j].charAt(sj) && split[j].charAt(sj) != '?') {
                        return false;
                    }

                    if (i < temp && sj != 0) {
                        return false;
                    }
                }
                return canNotMatch || i == temp - 1;
            }

            temp = matched? i : temp + 1;
            if (matched) {
                j++;
            }

            if (!canNotMatch && !matched) {
                return false;
            }

            matchedCount = matched ? matchedCount + 1 : matchedCount;
            canNotMatch = true;
        }

        return matchedCount == split.length && (canNotMatchLength || temp == text.length());
    }
    //private boolean matched = false;
    //
    //public boolean isMatch(String text, String pattern) {
    //    matched = false;
    //    match(0, 0, text, pattern);
    //    return matched;
    //}
    //
    //private void match(int textIndex, int patternIndex, String text, String pattern) {
    //    if (matched) {
    //        return;
    //    }
    //
    //    if (patternIndex == pattern.length() || textIndex == text.length()) {
    //        if (patternIndex == pattern.length()  ) {
    //            if (textIndex == text.length() || pattern.charAt(patternIndex - 1) == '*')
    //                matched = true;
    //        }
    //
    //        if (textIndex == text.length() && patternIndex < pattern.length()) {
    //            for (int i = patternIndex; i < pattern.length(); i++) {
    //                if (pattern.charAt(i) != '*') {
    //                    return;
    //                }
    //            }
    //            matched = true;
    //        }
    //        return;
    //    }
    //
    //    if (pattern.charAt(patternIndex) == '?') {
    //        match(textIndex + 1, patternIndex + 1, text, pattern);
    //    } else if (pattern.charAt(patternIndex) == '*') {
    //        for (int i = 0; i < text.length() - textIndex; i++) {
    //            match(textIndex + i, patternIndex + 1, text, pattern);
    //        }
    //    } else if (pattern.charAt(patternIndex) == text.charAt(textIndex)) {
    //        match(textIndex + 1, patternIndex + 1, text, pattern);
    //    }
    //}

    public static void main(String[] args) {
        Solution44 search = new Solution44();
        System.out.println(search.isMatch("aab", "?*a?"));
        System.out.println(search.isMatch("aa", "*"));
        System.out.println(search.isMatch("bbbaba", "bb**??"));
        System.out.println(search.isMatch("a", "ab*"));
        System.out.println(search.isMatch("mississippi", "m??*ss*?i*pi"));
    }
}
//leetcode submit region end(Prohibit modification and deletion)
