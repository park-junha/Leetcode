/*  Solution: Lookup Hash Map for each character of input
    Time Complexity:    O(n+1)  (n is length of str, +1 for negative check)
    Space Complexity:   O(1)  (constant memory)
 */

func myAtoi(str string) int {
  handleSpaces := true
  isNeg := false
  val := 0
  const INT_MAX = int(^uint32(0) >> 1)
  table := map[rune]int{
    '0': 0,
    '1': 1,
    '2': 2,
    '3': 3,
    '4': 4,
    '5': 5,
    '6': 6,
    '7': 7,
    '8': 8,
    '9': 9,
  }
  for _, char := range str {
    if char == ' ' {
      if handleSpaces == false {
        break
      }
    } else if char == '-' || char == '+' {
      if handleSpaces == false {
        break
      }
      if char == '-' {
        isNeg = true
      }
      handleSpaces = false
    } else {
      if digit, ok := table[char]; ok {
        if isNeg == false && (
          val > INT_MAX / 10 || (
            val == INT_MAX / 10 &&
            digit > INT_MAX % 10)) {
          return INT_MAX
        } else if isNeg == true && (
          val > INT_MAX / 10 || (
            val == INT_MAX / 10 &&
            digit > INT_MAX % 10 + 1)) {
          return -INT_MAX - 1
        } else {
          val = val * 10 + digit
        }
      } else {
        break
      }
      handleSpaces = false
    }
  }
  if isNeg == true {
    val *= -1
  }
  return val
}
