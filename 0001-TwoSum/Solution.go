/*  Solution: Hash Map
    Time complexity:    O(n)
    Space complexity:   O(n)
 */

func twoSum(nums []int, target int) []int {
  numsChecked := make(map[int]int)
  for index, value := range nums {
    complement := target - value
    if compIndex, ok := numsChecked[complement]; ok {
      return []int{compIndex, index}
    } else {
      numsChecked[value] = index
    }
  }
  return []int{-1,-1}
}
