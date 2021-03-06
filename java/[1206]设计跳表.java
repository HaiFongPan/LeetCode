package pri.leetcode.leetcode.editor.cn;
//不使用任何库函数，设计一个跳表。
//
// 跳表是在 O(log(n)) 时间内完成增加、删除、搜索操作的数据结构。跳表相比于树堆与红黑树，其功能与性能相当，并且跳表的代码长度相较下更短，其设计思想与链表相似。 
//
// 例如，一个跳表包含 [30, 40, 50, 60, 70, 90]，然后增加 80、45 到跳表中，以下图的方式操作： 
//
// 
//Artyom Kalinin [CC BY-SA 3.0], via Wikimedia Commons 
//
// 跳表中有很多层，每一层是一个短的链表。在第一层的作用下，增加、删除和搜索操作的时间复杂度不超过 O(n)。跳表的每一个操作的平均时间复杂度是 O(log(n))，空间复杂度是 O(n)。 
//
// 在本题中，你的设计应该要包含这些函数： 
//
// 
// bool search(int target) : 返回target是否存在于跳表中。 
// void add(int num): 插入一个元素到跳表。 
// bool erase(int num): 在跳表中删除一个值，如果 num 不存在，直接返回false. 如果存在多个 num ，将它们全部删除。 
// 
//
// 了解更多 : https://en.wikipedia.org/wiki/Skip_list 
//
// 注意，跳表中可能存在多个相同的值，你的代码需要处理这种情况。 
//
// 样例: 
//
// 
//Skiplist skiplist = new Skiplist();
//
//skiplist.add(1);
//skiplist.add(2);
//skiplist.add(3);
//skiplist.search(0);   // 返回 false
//skiplist.add(4);
//skiplist.search(1);   // 返回 true
//skiplist.erase(0);    // 返回 false，0 不在跳表中
//skiplist.erase(1);    // 返回 true
//skiplist.search(1);   // 返回 false，1 已被擦除
// 
//
// 约束条件: 
//
// 
// 0 <= num, target <= 20000 
// 最多调用 50000 次 search, add, 以及 erase操作。 
// 
// Related Topics 设计

import java.util.ArrayList;
import java.util.List;
import java.util.Random;

//leetcode submit region begin(Prohibit modification and deletion)
class Skiplist {
    private Node head;
    private int listLevel;
    private Random random = new Random();


    public Skiplist() {
        listLevel = 0; // 默认等级为 0
        head = new Node(-1);
    }
    
    public boolean search(int target) {
        Node search = head;
        int lvl = listLevel;

        while (lvl >= 0) {
            Node lvlNext = search.levelNext[lvl];
            // 下一个节点是 null 或者，下一个节点大于目标，需要下沉索引
            if (lvlNext == null || lvlNext.value > target) {
                lvl--;
            } else if (lvlNext.value < target) {
                search = lvlNext;
            } else {
                return true;
            }
        }

        return false;
    }
    
    public void add(int num) {
        Node search = head, target = null;
        int searchLevel = listLevel;

        int nodeLevel =  randomLevel(); // 计算添加节点的索引高度
        // 如果内部有相同的节点，则不增加其索引高度
        Node nNode = new Node(num);
        // 动态扩增索引
        if (nodeLevel > listLevel) {
            for (int i = nodeLevel; i > listLevel; i--) {
                head.levelNext[i] = nNode;
                nNode.levelPrev[i] = head;
            }

            listLevel = nodeLevel;
        }

        while (target == null) {
            final Node lvlNext = search.levelNext[searchLevel];
            if (lvlNext == null || lvlNext.value > num) {
                // 记录下沉节点
                if (searchLevel <= nodeLevel) {
                    search.levelNext[searchLevel] = nNode;
                    nNode.levelNext[searchLevel] = lvlNext;
                    nNode.levelPrev[searchLevel] = search;
                    if (lvlNext != null) {
                        lvlNext.levelPrev[searchLevel] = nNode;
                    }
                }
                if (searchLevel == 0) {
                    target = search;
                } else {
                    searchLevel--;
                }
            } else {
                search = lvlNext;
            }
        }
    }

    public boolean erase(int num) {
        Node search = head, target = null;
        int lvl = listLevel;

        while (lvl >= 0 && target == null) {
            final Node lvlNext = search.levelNext[lvl];
            if (lvlNext == null || lvlNext.value > num) {
                lvl--;
            } else if (lvlNext.value < num) {
                search = lvlNext;
            } else {
                target = lvlNext;
            }
        }

        if (target == null) {
            return false;
        }

        while (lvl >= 0) {
            if (target.levelNext[lvl] != null) {
                target.levelNext[lvl].levelPrev[lvl] = target.levelPrev[lvl];
            }
            target.levelPrev[lvl].levelNext[lvl] = target.levelNext[lvl];
            lvl--;
        }

        while (head.levelNext[listLevel] == null) {
            listLevel--;
        }

        return true;
    }

    // 随机决定增加节点的 levelNext
    private int randomLevel() {
        int lvl = 0;
        while (random.nextInt(2) > 0 && lvl < 15) {
            lvl++;
        }

        return lvl;
    }

    class Node {
        private int value;
        // 多级索引指针
        private Node[] levelNext;
        private Node[] levelPrev;

        Node(int value) {
            this.value = value;
            this.levelNext = new Node[20];
            this.levelPrev = new Node[20];
        }
    }

    public static void main(String[] args) {
        Skiplist skipList = new Skiplist();
        skipList.add(16);
        skipList.add(5);
        skipList.add(14);
        skipList.add(13);
        skipList.add(0);
        skipList.add(3);
        skipList.add(12);
        skipList.add(9);
        skipList.add(12);
        System.out.println(skipList.erase(3) );
        System.out.println(skipList.search(6) );
        skipList.add(7);
        System.out.println(skipList.erase(0) );
        System.out.println(skipList.erase(1) );
        System.out.println(skipList.erase(10) );
        skipList.add(5);
        System.out.println(skipList.search(12) );
        System.out.println(skipList.search(7) );
        System.out.println(skipList.search(16) );
        System.out.println(skipList.erase(7) );
        System.out.println(skipList.search(0) );
        skipList.add(9);
        skipList.add(16);
        skipList.add(3);
        System.out.println(skipList.erase(2) );
        System.out.println(skipList.search(17) );
        skipList.add(2);
        System.out.println(skipList.search(17) );
        System.out.println(skipList.erase(0) );
        System.out.println(skipList.search(9) );
        System.out.println(skipList.search(14) );
        System.out.println(skipList.erase(1) );
        System.out.println(skipList.erase(6) );
        skipList.add(1);
        System.out.println(skipList.erase(16) );
        System.out.println(skipList.search(9) );
        System.out.println(skipList.erase(10) );
        System.out.println(skipList.erase(9) );
        System.out.println(skipList.search(2) );
        skipList.add(3);
        skipList.add(16);
        System.out.println(skipList.erase(15) );
        System.out.println(skipList.erase(12) );
        System.out.println(skipList.erase(7) );
        skipList.add(4);
        System.out.println(skipList.erase(3) );
        skipList.add(2);
        System.out.println(skipList.erase(1) );
        System.out.println(skipList.erase(14) );
        skipList.add(13);
        skipList.add(12);
        skipList.add(3);
        System.out.println(skipList.search(6) );
        System.out.println(skipList.search(17) );
        skipList.add(2);
        System.out.println(skipList.erase(3) );
        System.out.println(skipList.search(14) );
        skipList.add(11);
        skipList.add(0);
        System.out.println(skipList.search(13) );
        skipList.add(2);
        System.out.println(skipList.search(1) );
        System.out.println(skipList.erase(10) );
        System.out.println(skipList.erase(17) );
        System.out.println(skipList.search(0) );
        System.out.println(skipList.search(5) );
        System.out.println(skipList.erase(8) );
        System.out.println(skipList.search(9) );
        skipList.add(8);
        System.out.println(skipList.erase(11) );
        System.out.println(skipList.search(10) );
        System.out.println(skipList.erase(11) );
        System.out.println(skipList.search(10) );
        System.out.println(skipList.erase(9) );
        System.out.println(skipList.erase(8) );
        System.out.println(skipList.search(15) );
        System.out.println(skipList.search(14) );
        skipList.add(1);
        skipList.add(6);
        skipList.add(17);
        skipList.add(16);
        System.out.println(skipList.search(13) );
        System.out.println(skipList.search(4) );
        System.out.println(skipList.search(5) );
        System.out.println(skipList.search(4) );
        System.out.println(skipList.search(17) );
        System.out.println(skipList.search(16) );
        System.out.println(skipList.search(7) );
        System.out.println(skipList.search(14) );
        System.out.println(skipList.search(1) );
    }
}

/**
 * Your Skiplist object will be instantiated and called as such:
 * Skiplist obj = new Skiplist();
 * boolean param_1 = obj.search(target);
 * obj.add(num);
 * boolean param_3 = obj.erase(num);
 */
//leetcode submit region end(Prohibit modification and deletion)
