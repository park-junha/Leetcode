/*  Solution: Two digits at a time
    Time Complexity:    O(ceil(n/2) + 1) (n is number of digits of x)
    Space Complexity:   O(1)
 */

import "math"

func isPalindrome(x int) bool {
  if x < 0 {
    return false
  }
  digits := log(x)
  for x > 0 {
    if x / int(math.Pow10(digits - 1)) != x % 10 {
      return false
    }
    x %= int(math.Pow10(digits - 1))
    x /= 10
    digits -= 2
  }
  return true
}

func log(n int) int {
  d := 0
  for n > 0 {
    n /= 10
    d++
  }
  return d
}
