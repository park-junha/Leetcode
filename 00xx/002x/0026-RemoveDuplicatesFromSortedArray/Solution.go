/*  Solution: Slicing
    Time Complexity:    O(n)
    Space Complexity:   O(1)
 */

func removeDuplicates(nums []int) int {
  i := 1
  for i < len(nums) {
    if nums[i] == nums[i-1] {
      nums = append(nums[:i], nums[i+1:]...)
    } else {
      i++
    }
  }
  return i
}
