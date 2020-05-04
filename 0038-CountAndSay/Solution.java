/*  Solution: Hash map
    Time complexity:    O(nk)
    Space complexity:   O(nk)

    Note: k is the length of the string, n the number of terms
 */

class Solution {
    public String countAndSay(int n) {
        //  Base case
        if (n == 1) {
            return "1";
        }
        
        //  Get previous term of sequence
        String prevTerm = countAndSay(n - 1);
        
        //  Initialize return value and hash map
        String newTerm = "";
        HashMap <String, Integer> map = new HashMap<> ();
        int temp;
        map.put("current", Integer.parseInt(
            String.valueOf(
                prevTerm.charAt(0)
            )
        ));
        map.put("occurrences", 0);
        for (int i = 0; i < prevTerm.length(); i++) {
            //  Count number of consecutive occurrences of a digit
            if (map.get("current") == Integer.parseInt(
                String.valueOf(
                    prevTerm.charAt(i)
                )
            )) {
                temp = map.get("occurrences");
                map.put("occurrences", temp + 1);
            }
            //  Append to return string once new digit appears
            else {
                newTerm += map.get("occurrences").toString();
                newTerm += map.get("current").toString();
                map.put("current", Integer.parseInt(
                    String.valueOf(
                        prevTerm.charAt(i)
                    )
                ));
                map.put("occurrences", 1);
            }
        }
        newTerm += map.get("occurrences").toString();
        newTerm += map.get("current").toString();
        return newTerm;
    }
}
