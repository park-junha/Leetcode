/*  Solution: Divide and Conquer with Recursion
    Time Complexity:    O(log n)
    Space Complexity:   O(log n)
 */

func searchInsert(nums []int, target int) int {
  i := len(nums) / 2
  for i >= 0 && i < len(nums) {
    if nums[i] == target {
      return i
    } else if target < nums[i] {
      return searchInsert(nums[:i], target)
    } else {
      return i + 1 + searchInsert(nums[i+1:], target)
    }
  }
  return 0
}
