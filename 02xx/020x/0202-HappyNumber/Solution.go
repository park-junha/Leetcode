/*  Solution: Hash Map
    Algorithm seems to be in NP? Not sure if in P (and therefore unsure of
    spacetime complexity
 */

func isHappy(n int) bool {
  seen := map[int]bool{n: true}
  for n != 1 {
    x := 0
    for n > 0 {
      x += (n % 10) * (n % 10)
      n /= 10
    }
    if _, ok := seen[x]; ok {
      return false
    }
    seen[x] = true
    n = x
  }
  return true
}
