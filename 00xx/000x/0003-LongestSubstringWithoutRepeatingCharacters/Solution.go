/*  Solution: Two Nested Loops
    Time Complexity:    O(n^2)
    Space Complexity:   O(n)

    Is there a better solution...?
 */

func lengthOfLongestSubstring(s string) int {
  max := 0
  for i, _ := range s {
    m := make(map[byte]bool)
    j := 0
    for j < len(s[i:]) {
      if _, ok := m[s[i+j]]; ok {
        break
      } else {
        m[s[i+j]] = true
      }
      j++
    }
    if j > max {
      max = j
    }
  }
  return max
}
