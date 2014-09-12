class ListNode {
    int val;
    ListNode next;
    ListNode(int x) {
         val = x;
         next = null;
    }
 }

public class Solution {
    public ListNode sortList(ListNode head) {
        if (head == null)
            return head;

        ListNode result = new ListNode(head.val);
        result.next = null;
        ListNode nextNode = head.next;

        while (nextNode != null) {
            ListNode pre = null;
            ListNode now = result;
            while (now != null && nextNode.val >= now.val){
                pre = now;
                now = now.next;
            }
            ListNode newNode = new ListNode(nextNode.val);
            if (now == null) {
                pre.next = newNode;
            }
            else if (pre==null) {
                newNode.next = now;
                result = newNode;
            } else {
                pre.next = newNode;
                newNode.next = now;
            }
            nextNode = nextNode.next;
        }

        return result;
    }

    // public static void main(String args[]){
    //     ListNode a = new ListNode(6);
    //     ListNode b = new ListNode(4);
    //     ListNode c = new ListNode(5);
    //     ListNode d = new ListNode(10);
    //     ListNode e = new ListNode(7);
    //     ListNode f = new ListNode(3);
    //     ListNode g = new ListNode(2);
    //     a.next = b;
    //     b.next = c;
    //     c.next = d;
    //     d.next = e;
    //     e.next = f;
    //     f.next = g;
    //     Solution s = new Solution();
    //     ListNode x = s.sortList(a);
    //     s.printtttt(x);
    // }

    // public void printtttt(ListNode x){
    //     while ( x != null) {
    //         System.out.print(x.val+"-->");
    //         x = x.next;
    //     }
    //     System.out.println();
    // }
}