package pri.leetcode.leetcode.editor.cn;
//给定一个区间的集合，找到需要移除区间的最小数量，使剩余区间互不重叠。
//
// 注意: 
//
// 
// 可以认为区间的终点总是大于它的起点。 
// 区间 [1,2] 和 [2,3] 的边界相互“接触”，但没有相互重叠。 
// 
//
// 示例 1: 
//
// 
//输入: [ [1,2], [2,3], [3,4], [1,3] ]
//
//输出: 1
//
//解释: 移除 [1,3] 后，剩下的区间没有重叠。
// 
//
// 示例 2: 
//
// 
//输入: [ [1,2], [1,2], [1,2] ]
//
//输出: 2
//
//解释: 你需要移除两个 [1,2] 来使剩下的区间没有重叠。
// 
//
// 示例 3: 
//
// 
//输入: [ [1,2], [2,3] ]
//
//输出: 0
//
//解释: 你不需要移除任何区间，因为它们已经是无重叠的了。
// 
// Related Topics 贪心算法

import java.util.Arrays;
import java.util.Comparator;

//leetcode submit region begin(Prohibit modification and deletion)
class Solution435 {
    public int eraseOverlapIntervals(int[][] intervals) {
        if (intervals.length <= 1) {
            return 0;
        }
        Arrays.sort(intervals, Comparator.comparingInt(o -> o[0]));
        //quickSort(intervals, 0, intervals.length - 1);

        int[] point = intervals[0];// min left
        int count = 1;
        for (int i = 1; i < intervals.length; i++) {
            if (point[1] <= intervals[i][0]) {
                // right <= next.left;
                count++;
                point = intervals[i];
            } else {
                if (point[1] >= intervals[i][1]) {
                    point = intervals[i];
                }
            }
        }

        return intervals.length - count;
    }

    private int split(int[][] array, int left, int right) {
        if (left >= right) {
            return left;
        }

        int[] pivot = array[right];
        int i = left;
        for (int j = left; j < right; j++) {
            // swap when
            if (array[j][0] < pivot[0]) {
                int[] temp = array[j];
                array[j] = array[i];
                array[i] = temp;
                i++;
            }
        }

        array[right] = array[i];
        array[i] = pivot;

        return i;
    }

    private void quickSort(int[][] array, int left, int right) {
        if (left < right) {
            final int pivot = split(array, left, right);
            quickSort(array, left, pivot - 1);
            quickSort(array, pivot + 1, right);
        }
    }

    public static void main(String[] args) {
        Solution435 solution435 = new Solution435();
        System.out.println(solution435.eraseOverlapIntervals(new int[][] {{1, 2}}));
        System.out.println(solution435.eraseOverlapIntervals(new int[][] {{1, 2}, {2, 3}, {3, 4}, {1, 3}}));
        System.out.println(solution435.eraseOverlapIntervals(new int[][] {{1, 2}, {1, 2}, {1, 2}}));
    }
}
//leetcode submit region end(Prohibit modification and deletion)
