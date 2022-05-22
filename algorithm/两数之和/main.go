package main

import "fmt"

// https://leetcode.cn/problems/two-sum/
// 解题思路: 查找 nums[i] + nums[j] == target
func twoSum(nums []int, target int) []int {
	// 先遍历 i
	for i := 0; i < len(nums); i++ {
		// 再遍历 j，j 永远大于 i，避免类似 (2,7) (7,2)这种重复组合
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func main() {
	nums := []int{2, 7, 8, 10}
	fmt.Println(twoSum(nums, 9))
}
