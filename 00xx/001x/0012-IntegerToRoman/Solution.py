# Solution: Hash Map
# Time Complexity:  O(nk)
# Space Complexity: O(n+k)

class Solution:
  def intToRoman(self, num: int) -> str:
    s = ""
    m = {
      1000: 'M'
      , 900: 'CM'
      , 500: 'D'
      , 400: 'CD'
      , 100: 'C'
      , 90: 'XC'
      , 50: 'L'
      , 40: 'XL'
      , 10: 'X'
      , 9: 'IX'
      , 5: 'V'
      , 4: 'IV'
      , 1: 'I'
    }
    
    while num > 0:
      for k in m:
        if num >= k:
          num -= k
          s += m[k]
          break
    
    return s
