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
}