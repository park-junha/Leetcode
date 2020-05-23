/*  Solution: Two Pointers + Character arithmetic
    Time Complexity:    O(n)
    Space Complexity:   O(1)

    Note: n is length of s.
          Also, could use regexp instead of character arithmetic
 */

func isPalindrome(s string) bool {
  i := 0
  j := len(s) - 1
  for i < j {
    for !isAlphaNum(s[i]) {
      i++
      if i >= j {
        return true
      }
    }
    for !isAlphaNum(s[j]) {
      j--
      if i >= j {
        return true
      }
    }
    if (s[i] - s[j]) % ('a' - 'A') != 0 {
      return false
    } else if s[i] != s[j] && (s[i] < 'A' || s[j] < 'A') {
      return false
    } else {
      i++
      j--
    }
  }
  return true
}

func isAlphaNum(x byte) bool {
  if x >= '0' && x <= '9' {
    return true
  } else if x >= 'A' && x <= 'Z' {
    return true
  } else if x >= 'a' && x <= 'z' {
    return true
  }
  return false
}
