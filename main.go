package main

import (
	easy_task "github.com/helga228/tasks/easy"
)

func main() {
	//int
	var numbers = []int{3, 9, 9}
	easy_task.IsPalindrome(121)
	easy_task.RomanToInt("III")
	easy_task.PlusOne(numbers)

	//string
	var strings = []string{"test", "test1", "test2", "re", "te"}
	easy_task.LongestCommonPrefix(strings)

}
