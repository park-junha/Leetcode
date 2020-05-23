/*  Solution: Brute Force
    Time Complexity:    O(n^2)
    Space Complexity:   O(1)
 */

func maxArea(height []int) int {
  max := 0
  for i := 0; i < len(height) - 1; i++ {
    for j := i + 1; j < len(height); j++ {
      var h int
      if height[i] < height[j] {
        h = height[i]
      } else {
        h = height[j]
      }
      if (j - i) * h > max {
        max = (j - i) * h
      }
    }
  }
  return max
}
