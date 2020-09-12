# Solution: Recursion
# Time Complexity:    O(3^n)
# Space Complexity:   O(n)

class Solution:
  def __init__(self):
    self.MODULO = 10 ** 9 + 7
    self.outputs = [
      [4,6],
      [6,8],
      [7,9],
      [4,8],
      [0,3,9],
      [],
      [1,7,0],
      [2,6],
      [1,3],
      [2,4]
    ]

  def getOutputs(self, n: int, start: int) -> int:
    if n == 1:
      return 1
    outputs = 0
    for nextStart in self.outputs[start]:
      outputs += self.getOutputs(n-1, nextStart)
    return outputs % self.MODULO

  def knightDialer(self, n: int) -> int:
    phoneNumbers = 0
    for start in range(10):
      phoneNumbers += self.getOutputs(n, start)
    return phoneNumbers % self.MODULO
