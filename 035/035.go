package main

import "fmt"

func searchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	temp := 0
	var flag bool
	for i, v := range nums {
		if v == target {
			return i
		}
		if target > v {
			temp = i
			flag = true
		}
		if target < v {
			temp = i
			flag = false
			break
		}
	}
	if flag {
		return temp + 1
	}
	return temp
}

func main() {
	var tests = []int{1, 3, 5, 6}
	fmt.Println(searchInsert(tests, 2))
}
