/*  Solution: Pop and push digits
    Time Complexity:    O(n) (or O(log_10(x))
    Space Complexity:   O(1)

    Note: n is the number of digits in x
 */

func reverse(x int) int {
  isNeg := false
  maxOne := int(^uint32(0) >> 1) % 10
  rev := 0
  if x < 0 {
    x = -x
    isNeg = true
    maxOne++
  }
  for x > 0 {
    digit := x % 10
    if (rev > int(^uint32(0) >> 1) / 10) || (rev == int(^uint32(0) >> 1) / 10 && digit > maxOne) {
      return 0
    }
    rev = rev * 10 + digit
    x /= 10
  }
  if isNeg {
    rev *= -1
  }
  return rev
}
