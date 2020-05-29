# Solution: Check previous index
# Time Complexity:  O(n)
# Space Complexity: O(1)

class Solution:
  def removeDuplicates(self, nums: List[int]) -> int:
    i = 1
    while i < len(nums):
      if nums[i-1] == nums[i]:
        del nums[i]
      else:
        i += 1
    return len(nums)
