# Solution: One Pass
# Time Complexity:    O(n)
# Space Complexity:   O(1)

class Solution:
  def distanceBetweenBusStops(self, distance: List[int], start: int, destination: int) -> int:
    if start > destination:
      start, destination = destination, start
    oneWay = sum(distance[start:destination])
    return min(oneWay, sum(distance) - oneWay)
