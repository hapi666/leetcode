package main

import "fmt"

func quckilySort(a []int) []int {
	if len(a) <= 1 {
		temp := make([]int, len(a))
		copy(temp, a)
		return temp
	}
	pivot := a[0]
	left := make([]int, 0, len(a))
	right := make([]int, 0, len(a))
	for index := 1; index < len(a); index++ {
		if a[index] > pivot {
			right = append(right, a[index])
		} else {
			left = append(left, a[index])
		}
	}
	right = quckilySort(right)
	left = quckilySort(left)
	return append(append(left, pivot), right...)
}

// 合并两个有序数组
func merge(nums1 []int, m int, nums2 []int, n int) {
	// temp := make([]int, m+n)
	for i, j := m, 0; j < n; i, j = i+1, j+1 {
		nums1[i] = nums2[j]
	}
	// nums1 = append(nums1[:m], nums2[:n]...)
	// copy(temp, quckilySort(nums1))
	// nums1 = quckilySort(nums1)
	copy(nums1, quckilySort(nums1))
	// for i, v := range nums1 {
	// 	fmt.Printf("index:%d\n%d\n", i, v)
	// }
}

func main() {
	n1 := []int{1, 2, 3, 0, 0, 0}
	n2 := []int{2, 5, 6}
	merge(n1, 3, n2, 3)
	for i, v := range n1 {
		fmt.Printf("index:%d\n%d\n", i, v)
	}
}
