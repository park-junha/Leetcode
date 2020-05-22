# Solution: Single Loop + Binomial Theorem
# Time Complexity:  O(n) (not counting nCk calculation)
# Space Complexity: O(1)

class Solution:
  def climbStairs(self, n: int) -> int:
    ways = 0
    ones = n
    twos = 0
    while ones >= 0:
      ways += self.nCk(ones + twos, twos)
      twos += 1
      ones -= 2
    return ways

  def nCk(self, n, k):
    nck = 1
    for i in range(k):
      nck *= (n - i)
      nck /= (i + 1)
    return int(nck)
