package main

import "fmt"

var (
	max = 1<<31 - 1
	min = -1 << 31
)

func isPalindrome(x int) bool {
	xx := x
	if x < 0 {
		return false
	}
	result := 0
	for x != 0 {
		result = result*10 + x%10
		x /= 10
	}
	if result > max || result < min {
		return false
	}
	if result == xx {
		return true
	}
	return false
}

func main() {
	test := 10
	fmt.Println(isPalindrome(test))
}
