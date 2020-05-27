# Solution: Stack
# Time Complexity: O(n)
# Space Complexity: O(n)

class Solution:
  def isValid(self, s: str) -> bool:
    pairs = {
      '}': '{',
      ']': '[',
      ')': '('
    }
    stack = []
    for c in s:
      if any(c == op for op in ['{','[','(']):
        stack.append(c)
      elif any(c == cl for cl in ['}',']',')']):
        if len(stack) == 0 or stack.pop() != pairs[c]:
          return False
    if len(stack) > 0:
      return False
    return True
