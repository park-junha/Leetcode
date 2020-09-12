# Solution: One Pass
# Time Complexity:    O(n)
# Space Complexity:   O(1)

class Solution:
  def distanceBetweenBusStops(self, distance: List[int], start: int, destination: int) -> int:
    forwardDistance = 0
    i = start
    while i != destination:
      forwardDistance += distance[i]
      if i + 1 >= len(distance):
        i = 0
      else:
        i += 1
    backwardDistance = 0
    while i != start:
      backwardDistance += distance[i]
      if i + 1 >= len(distance):
        i = 0
      else:
        i += 1
    if forwardDistance <= backwardDistance:
      return forwardDistance
    return backwardDistance
