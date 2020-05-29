# Solution: Binomial Theorem without Factorials
# Time Complexity:  O(T(n))
# Space Complexity: O(T(n))
#
# Note: T(n) is nth triangular number

class Solution:
  def generate(self, numRows: int) -> List[List[int]]:
    triangle = []
    for n in range(numRows):
      row = [1]
      for k in range(1, n+1):
        row.append(int(row[len(row) - 1] * (n-k+1) / k ) )
      triangle.append(row)
    return triangle
