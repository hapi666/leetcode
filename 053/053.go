package main

import "fmt"

// O(n^2)解法
func maxSubArray(nums []int) int {
	sum := 0
	max := nums[0]
	for i := range nums {
		sum = 0
		for index := i; index < len(nums); index++ {
			sum += nums[index]
			if sum > max {
				max = sum
			}
		}
	}
	return max
}

func main() {
	nums := []int{-2, -1}
	fmt.Println(maxSubArray(nums))
}
