package main

import "fmt"

// https://leetcode.cn/problems/two-sum/

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
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
