package main

import "fmt"

// O(n^2)解法
func maxSubArray1(nums []int) int {
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

// O(nlgn) 分治法 时间复杂度最小
func maxSubArray2(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	middle := len(nums) >> 1
	mleft := maxSubArray2(nums[:middle])
	mright := maxSubArray2(nums[middle:])
	maxl := nums[middle-1]
	suml := 0
	for index := middle - 1; index >= 0; index-- {
		suml += nums[index]
		if suml > maxl {
			maxl = suml
		}
	}
	maxr := nums[middle]
	sumr := 0
	for index := middle; index < len(nums); index++ {
		sumr += nums[index]
		if sumr > maxr {
			maxr = sumr
		}
	}
	max := maxl + maxr
	if max < mleft {
		max = mleft
	}
	if max < mright {
		max = mright
	}
	return max
}

// O(n) 动态规划算法
func maxSubArray3(nums []int) int {
	sum := nums[0]
	b := 0
	for index := 0; index < len(nums); index++ {
		if b > 0 {
			b += nums[index]
		} else {
			b = nums[index]
		}
		if b > sum {
			sum = b
		}
	}
	return sum
}

func main() {
	nums := []int{-2, -1}
	fmt.Println(maxSubArray3(nums))
}
