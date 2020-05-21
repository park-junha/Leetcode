/*  Solution: Hash Map
    Time Complexity:    O(n)
    Space Complexity:   O(n)
 */

func majorityElement(nums []int) int {
  m := make(map[int]int)
  majority := len(nums) / 2
  var value int
  for _, value = range nums {
    if _, ok := m[value]; ok {
      m[value]++
      if m[value] > majority {
        return value
      }
    } else {
      m[value] = 1
    }
  }
  return value
}
