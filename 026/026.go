package main

import "fmt"

func removeDuplicates(nums []int) int {
	fmt.Println("here!")
	for i := 0; i < len(nums); i++ {
		temp := nums[i]
		m := 1
		for _, n := range nums {
			if n == temp {
				m++
			}
		}
		nums = append(nums[:i+1], nums[i+m-1:]...)
	}
	// for _, v := range nums {
	// 	fmt.Print(v)
	// }
	// fmt.Println()
	return len(nums)
}

func main() {
	s := []int{1, 1, 2}
	fmt.Println(removeDuplicates(s))
}
