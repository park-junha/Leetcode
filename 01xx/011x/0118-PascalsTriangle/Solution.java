/*  Solution: 
    Time complexity: O(T(n))
    Space complexity: O(T(n))

    Note: T(n) is nth triangular number.
 */

class Solution {
    public List<List<Integer>> generate(int numRows) {
        LinkedList<List<Integer>> triangle = new LinkedList<List<Integer>>();
        for (int n = 0; n < numRows; n++) {
            LinkedList<Integer> currRow = new LinkedList<Integer>();
            //  Base case
            int nCk = 1;
            for (int k = 0; k <= n; k++) {
                currRow.add(nCk);
                //  Clever trick to calculate n choose k+1 without factorials
                nCk = nCk * (n - k) / (k + 1);
            }
            triangle.add(currRow);
        }
        return triangle;
    }
}
