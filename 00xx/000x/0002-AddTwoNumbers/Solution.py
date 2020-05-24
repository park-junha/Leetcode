# Solution: Elementary Math
# Time Complexity:    O(n+1)
# Space Complexity:   O(n+1)
#
# Note: n is the length of the larger list

# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
  def addTwoNumbers(self, l1: ListNode, l2: ListNode) -> ListNode:
    l3 = None
    carry = 0
    while l1 != None or l2 != None:
      op1 = 0
      op2 = 0
      if l1 != None:
        op1 = l1.val
        l1 = l1.next
      if l2 != None:
        op2 = l2.val
        l2 = l2.next
      
      digit = (op1 + op2 + carry) % 10
      
      if op1 + op2 + carry > 9:
        carry = 1
      else:
        carry = 0
      
      if l3 == None:
        l3 = ListNode(digit)
        curr = l3
      else:
        curr.next = ListNode(digit)
        curr = curr.next
    
    if carry == 1:
      curr.next = ListNode(1)
    return l3
