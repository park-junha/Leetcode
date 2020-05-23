/*  Solution: Hash Map
    Time Complexity:    O(nk)
    Space Complexity:   O(n)

    Note: k is number of elements in hash map. In this case, k=2
 */

class Solution {
  public List<String> fizzBuzz(int n) {
    List<String> ret = new LinkedList<String>();
    HashMap<Integer, String> map = new HashMap<Integer, String>();
    map.put(3, "Fizz");
    map.put(5, "Buzz");
    String term;
    for (int i = 1; i <= n; i++) {
      term = "";
      for (Map.Entry<Integer, String> entry : map.entrySet()) {
        if (i % entry.getKey() == 0) {
          term += entry.getValue();
        }
      }
      if (term == "") {
        ret.add(Integer.toString(i));
      }
      else {
        ret.add(term);
      }
    }
    return ret;
  }
}
