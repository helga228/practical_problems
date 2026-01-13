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

func SearchInsert(arr []int, target int) int {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}
