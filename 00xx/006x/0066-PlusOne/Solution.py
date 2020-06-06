# Solution: Right-to-left
# Time Complexity:  O(n)
# Space Complexity: O(1)

class Solution:
  def plusOne(self, digits: List[int]) -> List[int]:
    carry = 0
    for i in range(len(digits)-1, -1, -1):
      if digits[i] == 9:
        digits[i] = 0
      else:
        digits[i] += 1
        return digits
    return [1] + digits
