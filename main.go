package main

import (
	easy_task "github.com/helga228/tasks/easy"
)

func main() {
	//int
	easy_task.IsPalindrome(121)
	easy_task.RomanToInt("III")
	easy_task.PlusOne([]int{3, 9, 9})

	//string
	easy_task.LongestCommonPrefix([]string{"test", "test1", "test2", "re", "te"})
	easy_task.IsValid("[]{}({)")
	easy_task.LengthOfLastWord("text for test_func j ")

	//array
	easy_task.TwoSum([]int{6, 8, 6, 9}, 12)
	easy_task.SearchInsert([]int{1, 2, 3, 4, 5}, 5)
}
