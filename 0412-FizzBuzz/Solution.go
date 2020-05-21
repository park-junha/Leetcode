/*  Solution: Single Loop
    Time Complexity:    O(n)
    Space Complexity:   O(n) (O(1) if we are only counting extra space)
 */

import "strconv"

func fizzBuzz(n int) []string {
  s := []string{}
  var t string
  for i := 1; i <= n; i++ {
    t = ""
    if i % 3 == 0 {
      t += "Fizz"
    }
    if i % 5 == 0 {
      t += "Buzz"
    }
    if t == "" {
      s = append(s, strconv.Itoa(i))
    } else {
      s = append(s, t)
    }
  }
  return s
}
