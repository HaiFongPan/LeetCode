class ListNode {
    int val;
    ListNode next;
    ListNode(int x) {
         val = x;
         next = null;
    }
 }
/**
 * 普通归并排序
 */
public class Solution {
    public ListNode sortList(ListNode head) {
        if (head == null || head.next == null)
            return head;
        ListNode slow = head;
        ListNode fast = head;
        while (fast.next != null && fast.next.next != null){
            slow = slow.next;
            fast = fast.next.next;
        }

        ListNode L1 = head;
        ListNode L2 = slow.next;
        slow.next = null;
        ListNode result;
        ListNode r1 = sortList(L1);
        ListNode r2 = sortList(L2);
        if (r1.val <= r2.val) {
            result = r1;
            r1 = r1.next;
        } else {
            result = r2;
            r2 = r2.next;
        }

        ListNode headTemp = result;

        while (r1 != null && r2 != null){
            if (r1.val <= r2.val) {
                headTemp.next = r1;
                r1 = r1.next;
            } else {
                headTemp.next = r2;
                r2 = r2.next;
            }
            headTemp = headTemp.next;
        }

        if (r1 != null) {
            headTemp.next = r1;
        } else {
            headTemp.next = r2;
        }
        return result;
    }
}