# Solution: Two windows
# Time Complexity:    O(n)
# Space Complexity:   O(1)

class Solution:
  def isPalindrome(self, s: str) -> bool:
    i = 0
    j = len(s) - 1
    while i <= j:
      if s[i].isalnum() is False:
        i += 1
        continue
      if s[j].isalnum() is False:
        j -= 1
        continue
      if s[i].lower() != s[j].lower():
        return False
      i += 1
      j -= 1
    return True
