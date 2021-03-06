package pri.leetcode.leetcode.editor.cn;
//给定一个字符串 s 和一些长度相同的单词 words。找出 s 中恰好可以由 words 中所有单词串联形成的子串的起始位置。
//
// 注意子串要与 words 中的单词完全匹配，中间不能有其他字符，但不需要考虑 words 中单词串联的顺序。 
//
// 
//
// 示例 1： 
//
// 输入：
//  s = "barfoothefoobarman",
//  words = ["foo","bar"]
//输出：[0,9]
//解释：
//从索引 0 和 9 开始的子串分别是 "barfoo" 和 "foobar" 。
//输出的顺序不重要, [9,0] 也是有效答案。
// 
//
// 示例 2： 
//
// 输入：
//  s = "wordgoodgoodgoodbestword",
//  words = ["word","good","best","word"]
//输出：[]
// 
// Related Topics 哈希表 双指针 字符串

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

//leetcode submit region begin(Prohibit modification and deletion)
class Solution30 {
    public List<Integer> findSubstring(String s, String[] words) {
        if (words.length == 0 || words[0].length() == 0 || s.length() == 0) {
            return new ArrayList<>();
        }
        int step = words[0].length();
        List<Integer> result = new ArrayList<>();
        Map<String, Integer> map = new HashMap<>();
        for (String word: words) {
            map.put(word, map.getOrDefault(word, 0) + 1);
        }
        final int wordsLength = step * words.length;
        for (int i = 0; i <= s.length() - wordsLength; i++) {
            int count = 0;
            Map<String, Integer> temp = new HashMap<>();
            for (int j = i; j < i + wordsLength && j <= s.length() - step ; j+=step) {
                final String substring = s.substring(j, j + step);
                if (!map.containsKey(substring)) {
                    break; // 不存在词
                } else if (temp.getOrDefault(substring, 0).equals(map.get(substring))) {
                    break; // 词超了
                }
                // 记录出现次数
                temp.put(substring, temp.getOrDefault(substring, 0) + 1);
                count++; // 总词数增加
                if (count == words.length) {
                    result.add(i);
                }
            }
        }
        return result;
    }

    public static void main(String[] args) {
        String s = "barfoothefoobarman";
        String[] words = {"foo","bar"};

        Solution30 solution30 = new Solution30();
        final List<Integer> substring = solution30.findSubstring(s, words);
        System.out.println(substring);
    }
}
//leetcode submit region end(Prohibit modification and deletion)
