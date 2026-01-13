package easy_task

import "fmt"

func TwoSum(nums []int, target int) []int {
	start := 0
	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[start] == target {
			fmt.Println([]int{start, i})
			return []int{start, i}
		}
		if start < len(nums) && i == len(nums) {
			i = 0
			start++
		}
	}
	return []int{-1, -1}
}

func SearchInsert(nums []int, target int) int {
	med := len(nums) / 2
	if nums[med] == target {
		fmt.Println(med)
		return med
	}
	if nums[med] > target {
		SearchInsert(nums[:med], target)
	}
	if nums[med] < target {
		SearchInsert(nums[med:], target)
	}
	fmt.Println(med)
	return med
}
