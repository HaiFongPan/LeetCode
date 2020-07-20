package pri.leetcode.leetcode.editor.cn;
//给定一个字符串数组，将字母异位词组合在一起。字母异位词指字母相同，但排列不同的字符串。
//
// 示例: 
//
// 输入: ["eat", "tea", "tan", "ate", "nat", "bat"],
//输出:
//[
//  ["ate","eat","tea"],
//  ["nat","tan"],
//  ["bat"]
//] 
//
// 说明： 
//
// 
// 所有输入均为小写字母。 
// 不考虑答案输出的顺序。 
// 
// Related Topics 哈希表 字符串

import java.util.ArrayList;
import java.util.Arrays;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

//leetcode submit region begin(Prohibit modification and deletion)
class Solution49 {
    public List<List<String>> groupAnagrams(String[] strs) {
        Map<String, List<String>> dict = new HashMap<>();
        for (String str: strs) {
            String sort = sort(str);
            if (dict.containsKey(sort)) {
                dict.get(sort).add(str);
            } else {
                List<String> list = new ArrayList<>();
                list.add(str);
                dict.put(sort, list);
            }
        }
        return new ArrayList<>(dict.values());
    }

    private String sort(String str) {
        char[] arr = str.toCharArray();
        Arrays.sort(arr);
        return String.valueOf(arr);
    }

    public static void main(String[] args) {
        Solution49 solution49 = new Solution49();
        System.out.println(solution49.groupAnagrams(new String[]{"eat", "tea", "tan", "ate", "nat", "bat"}));
    }
}
//leetcode submit region end(Prohibit modification and deletion)
