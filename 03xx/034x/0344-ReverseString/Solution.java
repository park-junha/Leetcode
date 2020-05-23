/*  Solution: Two Pointers
    Time Complexity:    O(ceil(n/2))
    Space Complexity:   O(1)
 */

class Solution {
  public void reverseString(char[] s) {
    int i = 0;
    int j = s.length - 1;
    char temp;
    while (i < j) {
      temp = s[i];
      s[i] = s[j];
      s[j] = temp;
      i++;
      j--;
    }
  }
}
