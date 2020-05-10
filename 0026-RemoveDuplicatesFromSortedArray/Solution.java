/*  Solution: Hashmap + Single Pass Loop
    Time complexity:    O(n^2)
    Space complexity:   O(n)

    Note: This is a pretty bad solution
 */

class Solution {
    public int removeDuplicates(int[] nums) {
        HashMap appearedNums = new HashMap<Integer, Integer>();
        int currentLength = nums.length;
        for (int i = 0; i < currentLength; i++) {
            if (!appearedNums.containsKey(nums[i])) {
                appearedNums.put(nums[i], 1);
            }
            else {
                for (int j = i + 1; j < nums.length; j++) {
                    nums[j - 1] = nums[j];
                }
                currentLength--;
                i--;
            }
        }
        return appearedNums.size();
    }
}
