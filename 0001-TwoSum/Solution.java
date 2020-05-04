/*  Solution: Hash Map
    Time complexity:    O(n)
    Space complexity:   O(n)
 */

class Solution {
    public int[] twoSum(int[] nums, int target) {
        HashMap<Integer, Integer> map = new HashMap<>();
        int complement;
        for (int i = 0; i < nums.length; i++) {
            complement = target - nums[i];
            if (map.containsKey(complement)) {
                return new int[] {map.get(complement), i};
            }
            else {
                map.put(nums[i], i);
            }
        }
        return new int[] {};
    }
}
