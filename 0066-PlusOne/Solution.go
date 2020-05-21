/*  Solution: Single Loop
    Time Complexity:    O(n)
    Space Complexity:   O(n)
 */

func plusOne(digits []int) []int {
  carry := 1
  index := len(digits) - 1
  for carry > 0 || index >= 0 {
    if digits[index] == 9 {
      carry = 0
      digits[index] = 0
    } else {
      digits[index]++
      return digits
    }
    index--
  }
  return append([]int{1}, digits...)
}
