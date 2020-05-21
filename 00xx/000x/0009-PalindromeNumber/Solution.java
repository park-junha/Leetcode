/*  Solution: Start at both ends and go inward
    Time complexity: O(n/2 + 1)
    Space complexity: O(1)
 */

import java.lang.Math;

class Solution {
    public boolean isPalindrome(int x) {
        if (x < 0) { return false; }
        int digits = 0;
        int mostSigDig, leastSigDig;
        while (x >= Math.pow(10, digits)) {
            digits++;
        }
        while (digits > 0) {
            mostSigDig = (int) ( x / Math.pow(10, --digits) );
            x -= mostSigDig * Math.pow(10, digits);
            if (digits > 0) {
                leastSigDig = x % 10;
                x /= 10;
                digits--;
            }
            else {
                leastSigDig = mostSigDig;
            }
            if (mostSigDig != leastSigDig) {
                return false;
            }
        }
        return true;
    }
}
