/*  Solution: Splicing
    Time complexity:    O(n)
    Space complexity:   O(n)
 */

/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode() {}
 *     ListNode(int val) { this.val = val; }
 *     ListNode(int val, ListNode next) { this.val = val; this.next = next; }
 * }
 */

class Solution {
    public ListNode mergeTwoLists(ListNode l1, ListNode l2) {
        ListNode l3 = new ListNode();
        ListNode currentNode = l3;
        int nextValue;
        while (l1 != null || l2 != null) {
            if (l1 == null) {
                nextValue = l2.val;
                l2 = l2.next;
            }
            else if (l2 == null) {
                nextValue = l1.val;
                l1 = l1.next;
            }
            else {
                if (l1.val < l2.val) {
                    nextValue = l1.val;
                    l1 = l1.next;
                }
                else {
                    nextValue = l2.val;
                    l2 = l2.next;
                }
            }
            currentNode.next = new ListNode(nextValue);
            currentNode = currentNode.next;
        }
        return l3.next;
    }
}
