# Solution: Single Loop + Hash Map
# Time Complexity:  O(n)
# Space Complexity: O(n)

class Solution:
  def fizzBuzz(self, n: int) -> List[str]:
    l = []
    m = {
      3: 'Fizz'
      , 5: 'Buzz'
    }
    for i in range(1, n+1):
      term = ''
      for j in m:
        if i % j == 0:
          term += m[j]
      if term == '':
        term = str(i)
      l.append(term)
    return l
