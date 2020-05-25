/*  Solution: Brute Force
    Time Complexity:    O(n^2)
    Space Complexity:   O(1)

    Note: This is a pretty obvious but mediocre solution.
 */

class Solution {
public:
  vector<int> twoSum(vector<int>& nums, int target) {
    for (int i = 0; i < nums.size(); i++) {
      for (int j = i + 1; j < nums.size(); j++) {
        if (nums[i] + nums[j] == target) {
          return vector<int>{i, j};
        }
      }
    }
    return vector<int>{};
  }
};
