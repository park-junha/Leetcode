/*  Solution: Brute Force
    Time Complexity:    O(sqrt(x))
    Space Complexity:   O(1)
 */

func mySqrt(x int) int {
  i := 1
  for i*i <= x {
    i++
  }
  return i-1
}
