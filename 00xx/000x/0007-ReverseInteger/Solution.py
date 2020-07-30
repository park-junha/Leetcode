# Solution: Pop and push digits
# Time Complexity:    O(n) (or O(log_10(x))
# Space Complexity:   O(1)
#
# Note: n is the number of digits in x

class Solution:
  def reverse(self, x: int) -> int:
    MIN = -2 ** 31
    MAX = 2 ** 31 - 1
    
    r = 0
    isNeg = x < 0
    
    while abs(x) > 0:
      r = r * 10 + abs(x) % 10
      if (isNeg is True and r > -MIN) or (isNeg is False and r > MAX):
        return 0
      x /= 10
      x = int(x)
    
    if isNeg is True:
      return -r
    return r
