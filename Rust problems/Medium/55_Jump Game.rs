/*
You are given an integer array nums. You are initially positioned at the array's first index, and each element in the array represents your maximum jump length at that position.

Return true if you can reach the last index, or false otherwise.



Example 1:

Input: nums = [2,3,1,1,4]
Output: true
Explanation: Jump 1 step from index 0 to 1, then 3 steps to the last index.
Example 2:

Input: nums = [3,2,1,0,4]
Output: false
Explanation: You will always arrive at index 3 no matter what. Its maximum jump length is 0, which makes it impossible to reach the last index.


Constraints:

1 <= nums.length <= 104
0 <= nums[i] <= 105

*/


impl Solution {
    #[allow(dead_code)]
    pub fn can_jump(nums: Vec<i32>) -> bool {
        let n = nums.len();
        let mut mx = 0;

        for i in 0..n {
            if mx < i {
                return false;
            }
            mx = std::cmp::max(mx, i + (nums[i] as usize));
        }

        true
    }
}


fn main() {
    let nums = vec![2,3,1,1,4];
    let res = Solution::can_jump(nums);
    println!("{}", res);
}






