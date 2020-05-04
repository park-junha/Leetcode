/*  Solution: Digit-by-digit
    Time complexity:    O(n)
    Space complexity:   O(n)

    Note: n is the length of the larger of the two lists.
 */

class Solution {
    public ListNode addTwoNumbers(ListNode l1, ListNode l2) {
        ListNode sum = new ListNode();
        ListNode currentNode = sum;
        int op1 = 0;
        int op2 = 0;
        int carry = 0;
        while (true) {
            //  Operand 1
            if (l1 == null) {
                op1 = 0;
            }
            else {
                op1 = l1.val;
                l1 = l1.next;
            }
            
            //  Operand 2
            if (l2 == null) {
                op2 = 0;
            }
            else {
                op2 = l2.val;
                l2 = l2.next;
            }
            
            //  Actual addition
            currentNode.val = op1 + op2 + carry;
            
            //  Determine carry
            if (currentNode.val >= 10) {
                currentNode.val %= 10;
                carry = 1;
            }
            else {
                carry = 0;
            }
            
            //  Break if no more operand nodes left
            if (l1 == null && l2 == null) {
                break;
            }
            
            //  Otherwise move to next digit
            currentNode.next = new ListNode();
            currentNode = currentNode.next;
        }
        
        //  Add extra 1 if carry remains
        if (carry == 1) {
            currentNode.next = new ListNode(1);
        }
        return sum;
    }
}
