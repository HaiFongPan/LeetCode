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
}