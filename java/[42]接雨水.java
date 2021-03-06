package pri.leetcode.leetcode.editor.cn;
//给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
//
// 
//
// 上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。 感谢 Marcos 贡献此图。 
//
// 示例: 
//
// 输入: [0,1,0,2,1,0,1,3,2,1,2,1]
//输出: 6 
// Related Topics 栈 数组 双指针

//leetcode submit region begin(Prohibit modification and deletion)
class Solution42 {
    public int trap(int[] height) {
        int capacity = 0;
        // find first left side
        for (int left = 0; left < height.length - 2; left++) {
            if (height[left] == 0 || height[left + 1] >= height[left]) {
                continue;
            }
            int right = left + 2, found = left + 1;
            while (right < height.length) {
                if (height[right] > height[left]) { // higher than left
                    found = right;
                    break;
                } else if (height[right] > height[found]) {
                    found = right;
                }
                right++;
            }
            if (found != left + 1) {
                // calc capacity
                int maxHeight = Math.min(height[found], height[left]);
                for (int i = left + 1; i < found; i++) {
                    capacity = capacity + maxHeight - height[i];
                }
                left = found - 1;// incr in forloop
            }
        }

        return capacity;
    }

    public static void main(String[] args) {
        Solution42 solution42 = new Solution42();
        System.out.println(solution42.trap(new int[] {4, 2, 0, 3, 2, 5}));
        System.out.println(solution42.trap(new int[] {4, 2, 3}));
        System.out.println(solution42.trap(new int[] {0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}));
        System.out.println(solution42.trap(new int[] {5, 5, 1, 7, 1, 1, 5, 2, 7, 6}));
    }
}
//leetcode submit region end(Prohibit modification and deletion)
