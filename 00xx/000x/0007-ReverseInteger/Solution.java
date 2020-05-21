/*  Solution: Digit-by-digit
    Time complexity:    O(n)
    Space complexity:   O(1)

    Note: n is the number of digits of x.
 */

class Solution {
    public int reverse(int x) {
        int reversed = 0;
        int toReverse = (x < 0) ? -x: x;
        boolean isNegative = (x < 0) ? true : false;
        while (toReverse > 0) {
            //  Check if number is about to overflow MAX
            if (
            reversed > Integer.MAX_VALUE / 10 || (
                reversed == Integer.MAX_VALUE / 10 &&
                toReverse % 10 > Integer.MAX_VALUE % 10
                )
            ) {
                return 0;
            }
            
            //  Also check MIN overflow
            if (
            reversed < Integer.MIN_VALUE / 10 || (
                reversed == Integer.MIN_VALUE / 10 &&
                toReverse % 10 < Integer.MIN_VALUE % 10
                )
            ) {
                return 0;
            }
            
            //  Then do reverse operation
            reversed = reversed * 10 + (toReverse % 10);
            toReverse = (toReverse - toReverse % 10) / 10;
        }
        reversed = (isNegative) ? -reversed : reversed;
        return reversed;
    }
}
