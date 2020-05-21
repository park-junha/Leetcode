/*  Solution: Two Pointers
    Time Complexity: O(ceil(n/2))
    Space Complexity: O(1)
 */

func reverseString(s []byte)  {
  i := 0
  j := len(s) - 1
  for i < j {
    s[i], s[j] = s[j], s[i]
    i++
    j--
  }
}
