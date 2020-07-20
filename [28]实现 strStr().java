package pri.leetcode.leetcode.editor.cn;
//实现 strStr() 函数。
//
// 给定一个 haystack 字符串和一个 needle 字符串，在 haystack 字符串中找出 needle 字符串出现的第一个位置 (从0开始)。如果不存在，则返回 -1。 
//
// 示例 1: 
//
// 输入: haystack = "hello", needle = "ll"
//输出: 2
// 
//
// 示例 2: 
//
// 输入: haystack = "aaaaa", needle = "bba"
//输出: -1
// 
//
// 说明: 
//
// 当 needle 是空字符串时，我们应当返回什么值呢？这是一个在面试中很好的问题。 
//
// 对于本题而言，当 needle 是空字符串时我们应当返回 0 。这与C语言的 strstr() 以及 Java的 indexOf() 定义相符。 
// Related Topics 双指针 字符串



//leetcode submit region begin(Prohibit modification and deletion)
class Solution28 {
    public int strStr(String haystack, String needle) {
        // KMP Version
        if (needle == null || needle.length() == 0) {
            return -1;
        }

        final int[] next = generateNext(needle);
        int maxMatch = 0;
        for (int i = 0; i < haystack.length(); i++) {
            while (maxMatch > 0 && haystack.charAt(i) != needle.charAt(maxMatch)) {
                maxMatch = next[maxMatch - 1];
            }

            if (haystack.charAt(i) == needle.charAt(maxMatch)) {
                maxMatch++;
            }

            if (maxMatch == needle.length()) {
                return i - needle.length() + 1;
            }
        }

        return -1;
    }

    private int[] generateNext(String needle) {
        int[] next = new int[needle.length()];
        // 计算 next 的用的是动态规划，目标前缀和后缀的次长子串是最长字串的最长字串
        // next[i] 表示该串的最长字串的长度，next[i] - 1 表示最长字串的最后一个字符的下标
        int maxLength = 0;// 记录当前 needle[0,i] 的最长临时值
        for (int i = 1; i < needle.length(); i++) {
            // 当遇到不匹配的时候，已匹配部分的最长匹配，从这一位开始继续做匹配查询
            while (maxLength > 0 && needle.charAt(i) != needle.charAt(maxLength)) {
                maxLength = next[maxLength - 1];
            }
            // 如果当前值跟最长匹配字串的下一个值相等，则 maxLenght++
            if (needle.charAt(i) == needle.charAt(maxLength)) {
                maxLength++;
            }

            next[i] = maxLength;
        }

        return next;
    }

    public static void main(String[] args) {
        Solution28 solution28 = new Solution28();
        System.out.println(solution28.strStr("ababaabbbbababbaabaaabaabbaaaabbabaabbbbbbabbaabbabbbabbbbbaaabaababbbaabbbabbbaabbbbaaabbababbabbbabaaabbaabbabababbbaaaaaaababbabaababaabbbbaaabbbabb","abbabbbabaa"));
    }
}
//leetcode submit region end(Prohibit modification and deletion)
