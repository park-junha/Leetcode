/*  Solution: Modulo algorithm
    Time Complexity:    O(n)
    Space Complexity:   O(n)

    Note: n is length of s, and while there may be nested loops each index
          of s is visited only once, so time complexity is O(n)
 */

func convert(s string, numRows int) string {
  var mod int
  if numRows == 1 {
    mod = 1
  } else {
    mod = 2 * (numRows - 1)
  }
  zz := ""
  for i := 1; i <= numRows; i++ {
    j := i - 1
    for j < len(s) {
      zz += string(s[j])
      if i == 1 || i == numRows {
        j += mod
      } else if j % mod == i - 1{
        j += mod - 2 * (i - 1)
      } else {
        j += 2 * (i - 1)
      }
    }
  }
  return zz
}
