# Solution: Dictionary (Hash Map)
# Time Complexity:  O(n)
# Space Complexity: O(n)

class Solution:
  def twoSum(self, nums: List[int], target: int) -> List[int]:
    numsTracked = {}
    for index, value in enumerate(nums):
      complement = target - value
      if complement in numsTracked:
        return [numsTracked[complement], index]
      else:
        numsTracked[value] = index
