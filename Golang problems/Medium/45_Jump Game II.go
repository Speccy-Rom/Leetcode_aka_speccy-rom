/*

You are given a 0-indexed array of integers nums of length n. You are initially positioned at nums[0].

Each element nums[i] represents the maximum length of a forward jump from index i. In other words, if you are at nums[i], you can jump to any nums[i + j] where:

0 <= j <= nums[i] and
i + j < n
Return the minimum number of jumps to reach nums[n - 1]. The test cases are generated such that you can reach nums[n - 1].



Example 1:

Input: nums = [2,3,1,1,4]
Output: 2
Explanation: The minimum number of jumps to reach the last index is 2. Jump 1 step from index 0 to 1, then 3 steps to the last index.
Example 2:

Input: nums = [2,3,0,1,4]
Output: 2


Constraints:

1 <= nums.length <= 104
0 <= nums[i] <= 1000
*/

package main

import "fmt"

func jump(nums []int) int {
	jumps := 0
	curEnd := 0
	curFarthest := 0
	for i := 0; i < len(nums)-1; i++ {
		curFarthest = Max(curFarthest, i+nums[i])
		if i == curEnd {
			jumps++
			curEnd = curFarthest
		}
	}
	return jumps

}

//func jump(nums []int) int {
//    res := 0
//    l := 0
//    r := 0
//    for r < len(nums) - 1 {
//        farthest := 0
//        for i := l; i <= r; i++ {
//            if i+ nums[i] > farthest {
//                farthest = i + nums[i]
//            }
//        }
//        l = r + 1
//        r = farthest
//        res++
//    }
//    return res
//}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	nums := []int{2, 3, 1, 1, 4}
	fmt.Println(jump(nums))
}
