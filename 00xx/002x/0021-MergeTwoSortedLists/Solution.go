/*  Solution: Splice input lists and return new list
    Time Complexity:    O(n)
    Space Complexity:   O(n)
 */

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
  l3 := &ListNode{}
  curr := l3
  for l1 != nil || l2 != nil {
    if l1 == nil {
      curr.Next = &ListNode{Val: l2.Val}
      curr = curr.Next
      l2 = l2.Next
    } else if l2 == nil {
      curr.Next = &ListNode{Val: l1.Val}
      curr = curr.Next
      l1 = l1.Next
    } else {
      if l1.Val < l2.Val {
        curr.Next = &ListNode{Val: l1.Val}
        curr = curr.Next
        l1 = l1.Next
      } else {
        curr.Next = &ListNode{Val: l2.Val}
        curr = curr.Next
        l2 = l2.Next
      }
    }
  }
  return l3.Next
}
