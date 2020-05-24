/*  Solution: Vertical scanning
    Time Complexity:    O(n*m)
    Space Complexity:   O(1)

    Note: n = number of strings
          m = length of longest string
 */

func longestCommonPrefix(strs []string) string {
  if len(strs) == 0 {
    return ""
  }
  for i, char := range strs[0] {
    for _, str := range strs[1:] {
      if i >= len(str) {
        return strs[0][:i]
      } else if rune(str[i]) != char {
        return strs[0][:i]
      }
    }
  }
  return strs[0]
}
