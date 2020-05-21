/*  Solution: Character-by-character
    Time complexity:    O(n^3)
    Space complexity:   O(1)
 */

class Solution {    
    public int myAtoi(String str) {
        boolean startedInt = false;
        boolean isNegative = false;
        int val = 0;
        for (int i = 0; i < str.length(); i++) {
            //  Case for digits 0-9
            if (str.charAt(i) == '0' ||
                str.charAt(i) == '1' ||
                str.charAt(i) == '2' ||
                str.charAt(i) == '3' ||
                str.charAt(i) == '4' ||
                str.charAt(i) == '5' ||
                str.charAt(i) == '6' ||
                str.charAt(i) == '7' ||
                str.charAt(i) == '8' ||
                str.charAt(i) == '9'
            ) {
                startedInt = true;
                int digit = 0;
                switch (str.charAt(i)) {
                    case '0':
                        digit = 0;
                        break;
                    case '1':
                        digit = 1;
                        break;
                    case '2':
                        digit = 2;
                        break;
                    case '3':
                        digit = 3;
                        break;
                    case '4':
                        digit = 4;
                        break;
                    case '5':
                        digit = 5;
                        break;
                    case '6':
                        digit = 6;
                        break;
                    case '7':
                        digit = 7;
                        break;
                    case '8':
                        digit = 8;
                        break;
                    case '9':
                        digit = 9;
                        break;
                }
                //  Check if return value is about to overflow
                int bound = (isNegative ? Integer.MIN_VALUE % 10 * -1 :
                             Integer.MAX_VALUE % 10);
                if (val > Integer.MAX_VALUE / 10 || (
                    val == Integer.MAX_VALUE / 10 &&
                    digit > bound
                )) {
                    return (
                        isNegative ?
                        Integer.MIN_VALUE :
                        Integer.MAX_VALUE
                    );
                }
                else {
                    val = val * 10 + digit;
                }
            }
            //  Handle +/- and toggle isNegative flag accordingly
            else if (str.charAt(i) == '+' || str.charAt(i) == '-') {
                if (startedInt) {
                    return isNegative ? -val : val;
                }
                else {
                    isNegative = (str.charAt(i) == '-') ? true : false;
                    startedInt = true;
                }
            }
            //  Handle pre-integer whitespaces
            else if (str.charAt(i) == ' ') {
                if (startedInt) {
                    return isNegative ? -val : val;
                }
                else {
                    continue;
                }
            }
            else {
                return isNegative ? -val : val;
            }
        }
        return isNegative ? -val : val;
    }
}
