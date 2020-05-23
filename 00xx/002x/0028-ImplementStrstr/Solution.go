/*  Solution: Check substring at each index
    Time Complexity:    O(n) (O(n^2) is substring operation counts as O(n))
    Space Complexity:   O(1)
 */

func strStr(haystack string, needle string) int {
  for i := 0; i <= len(haystack) - len(needle); i++ {
    if haystack[i:i + len(needle)] == needle {
      return i
    }
  }
  return -1
}
