/*  Solution: Elementary Math
    Time Complexity:    O(n)
    Space Complexity:   O(n)

    Note: n is the length of the larger of the two lists
 */

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
  l3 := &ListNode{}
  current := l3
  carry := 0
  var op1 int
  var op2 int
  for l1 != nil || l2 != nil {
    if l1 == nil {
      op1 = 0
      op2 = l2.Val
      l2 = l2.Next
    } else if l2 == nil {
      op1 = l1.Val
      op2 = 0
      l1 = l1.Next
    } else {
      op1 = l1.Val
      op2 = l2.Val
      l1 = l1.Next
      l2 = l2.Next
    }
    if carry + op1 + op2 < 10 {
      current.Val = carry + op1 + op2
      carry = 0
    } else {
      current.Val = (carry + op1 + op2) % 10
      carry = 1
    }
    if l1 != nil || l2 != nil {
      current.Next = &ListNode{}
      current = current.Next
    }
  }
  if carry == 1 {
    current.Next = &ListNode{Val: 1}
  }
  return l3
}
