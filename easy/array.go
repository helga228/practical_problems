package easy_task

import (
	"fmt"
	"sort"
)

func EasyBinarySearch(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			fmt.Println(mid)
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		}
	}
	return -1
}

func SearchInsertPosition(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			fmt.Println(mid)
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		}
	}
	fmt.Println(left)
	return left
}

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

func MajorityElement(nums []int) int {
	sorted := make([]int, len(nums))
	copy(sorted, nums)
	sort.Ints(sorted)

	fmt.Println(sorted[len(sorted)/2])
	return sorted[len(sorted)/2]
}
