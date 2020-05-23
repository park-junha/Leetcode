/*  Solution: Right-to-left
    Time Complexity:    O(n)
    Space Complexity:   O(1)

    Note: n is the length of s
 */

func lengthOfLastWord(s string) int {
  l := 0
  for i := len(s) - 1; i >= 0; i-- {
    if s[i] == ' ' {
      if l != 0 {
        return l
      }
    } else {
      l++
    }
  }
  return l
}
