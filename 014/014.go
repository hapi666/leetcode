package main

import (
	"fmt"
)

func longestCommonPrefix(strs []string) string {

	strsm := findMin(strs)
	index := 0
Label:
	for index = 0; index < strsm; index++ {
		temp := strs[0][index]
		ii := 0
		for ii < len(strs) {
			if strs[ii][index] != temp {
				break Label
			}
			ii++
		}
	}
	if index < 1 {
		return ""
	}
	return strs[0][:index]
}
func findMin(s []string) int {
	if len(s) == 0 {
		return 0
	}
	min := len(s[0])
	for _, v := range s {
		if min > len(v) {
			min = len(v)
		}
	}
	return min
}

func main() {
	test := []string{"flower", "flow", "flight"}
	tes := []string{}
	fmt.Println(longestCommonPrefix(test))
	fmt.Println(longestCommonPrefix(tes))
}
