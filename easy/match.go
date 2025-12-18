package easy_task

import "fmt"

func IsPalindrome(x int) bool {
	oldInt := x
	newInt := 0
	for x > 0 {
		digitReverse := x % 10
		newInt = newInt*10 + digitReverse
		x = x / 10
	}
	if newInt == oldInt {
		fmt.Printf("%d is palindrome\n", oldInt)
		return true
	}
	fmt.Printf("%d is not palindrome\n", oldInt)
	return false
}

func RomanToInt(s string) int {
	values := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	total := 0

	for i := 0; i < len(s); i++ {
		current := values[s[i]]
		if i+1 > len(s) && values[s[i+1]] > current {
			total -= current
		}
		total += current
	}
	fmt.Println(total)
	return total
}

func PlusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i]+1 == 10 {
			digits[i] = 0
		} else {
			digits[i] += 1
			break
		}
	}
	fmt.Println(digits)
	return digits
}
