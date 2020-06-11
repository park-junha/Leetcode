# Solution: Log / Modulo (no string conversion)
# Time Complexity:  O(log(n)) (n is x)
# Space Complexity: O(1)

import math

class Solution:
  def isPalindrome(self, x: int) -> bool:
    if x < 0:
      return False
    elif x == 0:
      return True
    d = int(math.log10(x)) + 1
    while d > 0:
      if int(x / math.pow(10, d-1)) != x % 10:
        return False
      x %= math.pow(10, d-1)
      x = int(x/10)
      d -= 2
    return True
