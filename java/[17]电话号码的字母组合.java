package pri.leetcode.leetcode.editor.cn;
//给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。
//
// 给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。 
//
// 
//
// 示例: 
//
// 输入："23"
//输出：["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
// 
//
// 说明: 
//尽管上面的答案是按字典序排列的，但是你可以任意选择答案输出的顺序。 
// Related Topics 字符串 回溯算法

import java.util.ArrayList;
import java.util.List;

//leetcode submit region begin(Prohibit modification and deletion)
class Solution17 {

    public List<String> letterCombinations(String digits) {
        char[][] c = new char[9][4];
        int letter = 0;
        for (int i = 0; i < 9; i++) {
            for (int j = 0; j < 4; j++) {
                if (i == 0) {
                    continue;
                } else if (i != 6 && i != 8 && j == 3) {
                    continue;
                }

                c[i][j] = (char) ('a' + letter);
                letter++;
            }
        }

        List<String> result = new ArrayList<>();
        for (int i = 0; i < digits.length(); i++) {
            result = combine(result, c, digits.charAt(i));
        }

        return result;
    }

    private List<String> combine(List<String> prev, char[][] c, char number) {
        int index = number - '1';
        List<String> combine = new ArrayList<>();
        if (prev.size() == 0) {
            for (int i = 0; i < 4; i++) {
                if (c[index][i] != 0) {
                    combine.add("" + c[index][i]);
                }
            }
        }
        for (String p : prev) {
            for (int i = 0; i < 4; i++) {
                if (c[index][i] != 0) {
                    combine.add(p + c[index][i]);
                }
            }
        }
        return combine;
    }

    public static void main(String[] args) {
        Solution17 solution17 = new Solution17();
        System.out.println(solution17.letterCombinations("23"));
    }
}
//leetcode submit region end(Prohibit modification and deletion)
