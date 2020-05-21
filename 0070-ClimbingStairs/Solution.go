/*  Solution: Binomial theorem + Single Pass
    Time Complexity:    O(n) (O(n^2 if we take nCk into account)
    Space Complexity:   O(n)
 */

func climbStairs(n int) int {
  ways := 0
  ones := n
  twos := 0
  for ones >= 0 {
    ways += nCk(ones + twos, twos)
    twos++
    ones -= 2
  }
  return ways
}

func nCk(n int, k int) int {
  nck := 1
  for i := 0; i < k; i++ {
    nck *= n - i
    nck /= i + 1
  }
  return nck
}
