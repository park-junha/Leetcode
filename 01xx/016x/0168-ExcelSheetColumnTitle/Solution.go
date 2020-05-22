/*  Solution: Character arithmetic
    Time Complexity:    O(log_26(n))
    Space Complexity:   O(log_26(n))
 */

func convertToTitle(n int) string {
  title := ""
  for n > 0 {
    title = string(rune((n - 1) % 26 + int('A'))) + title
    n -= (n - 1) % 26 + 1
    n /= 26
  }
  return title
}
