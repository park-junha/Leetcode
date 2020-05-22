/*  Solution: Character arithmetic
    Time Complexity:    O(n)
    Space Complexity:   O(1)
 */

func titleToNumber(s string) int {
  col := 0
  for _, char := range s {
    col = col * 26 + (int(char) - int('A') + 1)
  }
  return col
}
