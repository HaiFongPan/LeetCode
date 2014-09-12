class ListNode {
    int val;
    ListNode next;
    ListNode(int x) {
         val = x;
         next = null;
    }
 }

public class Solution {
    public ListNode reorderList(ListNode head) {
        if (head == null) {
            return head;
        }
        ListNode slow = head;
        ListNode fast = head;
        //split into two list by the middle one
        while (fast!= null && fast.next!=null) {
            slow = slow.next;
            fast = fast.next.next;
        }

        ListNode L1 = head;
        ListNode L2 = slow.next;
        slow.next = null;

        if (L2 == null) {
            return L1;
        }
        if (L2.next == null){
            L2.next = L1.next;
            L1.next = L2;
        } else {
            ListNode pointL1 = L1;
            ListNode pointL2 = L2;
            while (pointL2!= null) {
                pointL2 = L2;
                while (pointL2.next!=null && pointL2.next.next!=null)
                    pointL2 = pointL2.next;

                if (pointL2.next != null) {
                    ListNode tempNode = pointL2.next;
                    pointL2.next = null;
                    tempNode.next = pointL1.next;
                    pointL1.next = tempNode;
                } else if (pointL2 != null) {
                    ListNode tempNode = pointL2;
                    pointL2 = null;
                    tempNode.next = pointL1.next;
                    pointL1.next = tempNode;
                }

                pointL1 = pointL1.next.next;
            }
        }
        return L1;
    }

//     public static void main(String args[]){
//         ListNode a = new ListNode(6);
//         ListNode b = new ListNode(4);
//         ListNode c = new ListNode(5);
//         ListNode d = new ListNode(10);
//         ListNode e = new ListNode(7);
//         ListNode f = new ListNode(3);
//         ListNode g = new ListNode(2);
//         a.next = b;
//         b.next = c;
// //        c.next = d;
// //        d.next = e;
// //        e.next = f;
// //        f.next = g;
//         Solution s = new Solution();
//         ListNode x = s.reorderList(null);
//         s.printtttt(x);
//     }

//     public void printtttt(ListNode x){
//         while ( x != null) {
//             System.out.print(x.val+"-->");
//             x = x.next;
//         }
//         System.out.println();
//     }
}