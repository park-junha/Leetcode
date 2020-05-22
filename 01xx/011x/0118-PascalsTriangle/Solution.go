/*  Solution: Calculate nCk without factorials / incrementally
    Time Complexity:    O(n^2) / O(T(n))
    Space Complexity:   O(n^2) / O(T(n))

    Note: T(n) is nth triangular number
 */

func generate(numRows int) [][]int {
  triangle := [][]int{}
  for i := 0; i < numRows; i++ {
    row := []int{1}
    for j := 1; j < i + 1; j++ {
      row = append(row, row[len(row) - 1] * (i - j + 1) / j)
    }
    triangle = append(triangle, row)
  }
  return triangle
}
