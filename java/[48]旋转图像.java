package pri.leetcode.leetcode.editor.cn;
//给定一个 n × n 的二维矩阵表示一个图像。
//
// 将图像顺时针旋转 90 度。 
//
// 说明： 
//
// 你必须在原地旋转图像，这意味着你需要直接修改输入的二维矩阵。请不要使用另一个矩阵来旋转图像。
//
// 示例 1: 
//
// 给定 matrix = 
//[
//  [1,2,3],
//  [4,5,6],
//  [7,8,9]
//],
//
//原地旋转输入矩阵，使其变为:
//[
//  [7,4,1],
//  [8,5,2],
//  [9,6,3]
//]
// 
//
// 示例 2: 
//
// 给定 matrix =
//[
//  [ 5, 1, 9,11],
//  [ 2, 4, 8,10],
//  [13, 3, 6, 7],
//  [15,14,12,16]
//], 
//
//原地旋转输入矩阵，使其变为:
//[
//  [15,13, 2, 5],
//  [14, 3, 4, 1],
//  [12, 6, 8, 9],
//  [16, 7,10,11]
//]
// 
// Related Topics 数组

//leetcode submit region begin(Prohibit modification and deletion)
class Solution48 {
    public void rotate(int[][] matrix) {
        int n = matrix.length / 2;
        for (int i = 0; i < n; i++) {
            int len = matrix.length - 2 * i - 1; // if length = 7 , then i = 0, length = 6, i = 1, length = 4
            int magic = matrix.length - i - 1;
            for (int upj = i; upj < magic; upj++) {
                // swap martrix[i][upj] with coloum
                int targetI = i + (upj + len) % magic;
                int temp = matrix[targetI][magic];
                matrix[targetI][magic] = matrix[i][upj];
                matrix[i][upj] = temp;
            }
            for (int downj = magic; downj > i; downj--) {
                // matrix[row][downj] with coloum
                int targetI = magic + (downj - len) - i;
                int temp = matrix[targetI][i];
                matrix[targetI][i] = matrix[magic][downj];
                matrix[magic][downj] = temp;
            }

            for (int j1 = i, j2 = magic; j1 < magic; j1++) {
                int temp = matrix[i][j1];
                matrix[i][j1] = matrix[magic][j2];
                matrix[magic][j2] = temp;
                j2--;
            }
        }

        //print(matrix);
    }

    private void print(int[][] matrix) {
        for (int i = 0; i < matrix.length; i++) {
            for (int j = 0; j < matrix.length; j++) {
                System.out.print(matrix[i][j] + ",");
            }
            System.out.println();
        }
    }

    public static void main(String[] args) {
        Solution48 solution48 = new Solution48();
        solution48.rotate(new int[][] {
                {5, 1, 9, 11},
                {2, 4, 8, 10},
                {13, 3, 6, 7},
                {15, 14, 12, 16}
        });

        solution48.rotate(new int[][] {
                {7,4,1},
                {8,5,2},
                {9,6,3}
        });


    }
}
//leetcode submit region end(Prohibit modification and deletion)
