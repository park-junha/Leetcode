/*  Solution: Roman to Integer
    Time Complexity: O(n)
    Space Complexity: O(1)
 */

func romanToInt(s string) int {
  prefix := '0'
  val := 0
  for i, _ := range s {
    switch s[i] {
    case 'I':
      prefix = 'I'
      val += 1
    case 'V':
      if prefix == 'I' {
        prefix = '0'
        val += 3
      } else {
        prefix = 'V'
        val += 5
      }
    case 'X':
      if prefix == 'I' {
        prefix = '0'
        val += 8
      } else {
        prefix = 'X'
        val += 10
      }
    case 'L':
      if prefix == 'X' {
        prefix = '0'
        val += 30
      } else {
        prefix = 'L'
        val += 50
      }
    case 'C':
      if prefix == 'X' {
        prefix = '0'
        val += 80
      } else {
        prefix = 'C'
        val += 100
      }
    case 'D':
      if prefix == 'C' {
        prefix = '0'
        val += 300
      } else {
        prefix = 'D'
        val += 500
      }
    case 'M':
      if prefix == 'C' {
        prefix = '0'
        val += 800
      } else {
        prefix = 'M'
        val += 1000
      }
    }
  }
  return val
}
