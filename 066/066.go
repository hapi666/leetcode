package main

import "fmt"

func plusOne(digits []int) []int {
	if digits[0] == 0 {
		digits[0] = 1
		return digits
	}
	v := 1
	for index := len(digits) - 1; index >= 0; index-- {
		if digits[index]+v >= 10 {
			v = (digits[index] + v) / 10
			digits[index] = (digits[index] + v) % 10
			if index == 0 {
				digits = append([]int{v}, digits...)
			}
		} else {
			digits[index] += v
			v = 0
		}
	}
	return digits
}

func main() {
	test := []int{1, 8}
	result := plusOne(test)
	for _, v := range result {
		fmt.Println(v)
	}

}
