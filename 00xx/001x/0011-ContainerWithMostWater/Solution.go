/*  Solution: Two Windows
    Time Complexity:    O(n)
    Space Complexity:   O(1)

    Note: Keep track of two pointers and move the pointer at the shorter
          edge of the rectangle inwards.

          Moving the pointer of the longer edge will never yield a larger
          area because the shorter edge won't be altered + we decrease the
          width of the rectangle each time we move a pointer inwards
 */

func maxArea(height []int) int {
  max := 0
  i := 0
  j := len(height) - 1
  for i < j {
    var h int
    if height[i] < height[j] {
      h = height[i]
      if (j - i) * h > max {
        max = (j - i) * h
      }
      i++
    } else {
      h = height[j]
      if (j - i) * h > max {
        max = (j - i) * h
      }
      j--
    }
  }
  return max
}
