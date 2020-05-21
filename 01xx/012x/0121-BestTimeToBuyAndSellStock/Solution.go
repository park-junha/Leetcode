/*  Solution: Nested Loops (Brute Force)
    Time Complexity:    O(n^2) (more specifically, n(n-1)/2 iterations)
    Space Complexity:   O(1)

    Note: Not the best solution
 */

func maxProfit(prices []int) int {
  max := 0
  for i := 0; i < len(prices); i++ {
    for j := i + 1; j < len(prices); j++ {
      if prices[j] - prices[i] > max {
        max = prices[j] - prices[i]
      }
    }
  }
  return max
}
