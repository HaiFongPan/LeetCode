package pri.leetcode.leetcode.editor.cn;
//给出一个区间的集合，请合并所有重叠的区间。
//
// 示例 1: 
//
// 输入: [[1,3],[2,6],[8,10],[15,18]]
//输出: [[1,6],[8,10],[15,18]]
//解释: 区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
// 
//
// 示例 2: 
//
// 输入: [[1,4],[4,5]]
//输出: [[1,5]]
//解释: 区间 [1,4] 和 [4,5] 可被视为重叠区间。 
// Related Topics 排序 数组



//leetcode submit region begin(Prohibit modification and deletion)
class Solution56 {
    public int[][] merge(int[][] intervals) {
        if (intervals.length <= 1) {
            return intervals;
        }

        quickSort(intervals, 0, intervals.length - 1);
        int length = 0;
        int[][] temp = new int[intervals.length][2];
        int r = 0;
        for (int i = 0 ; i < intervals.length; i++) {
            if (i+1 == intervals.length || intervals[i][1] < intervals[i+1][0]) {
                // 无重叠
                temp[r++] = intervals[i];
                length++;
                continue;
            }

            intervals[i+1][0] = Math.min(intervals[i][0], intervals[i+1][0]);
            intervals[i+1][1] = Math.max(intervals[i][1], intervals[i+1][1]);
        }

        int[][] result = new int[length][2];
        System.arraycopy(temp, 0, result, 0, length);

        return result;
    }

    public int split(int[][] array, int left, int right) {
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

    public void quickSort(int[][] array, int left, int right) {
        if (left < right) {
            final int pivot = split(array, left, right);
            quickSort(array, left, pivot - 1);
            quickSort(array, pivot + 1, right);
        }
    }

    public static void main(String[] args) {
        Solution56 s = new Solution56();
        int[][] intervals = new int[][]{{1,4},{4,5}};
        final int[][] merge = s.merge(intervals);
        for (int[] m: merge) {
            System.out.println(String.format("[%d,%d]", m[0], m[1]));
        }
    }
}
//leetcode submit region end(Prohibit modification and deletion)
