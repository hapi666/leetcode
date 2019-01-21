package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	if len(strs) < 1 {
		return ""
	}
	preIndexLength := len(strs[0])
	for i, v := range strs {
		if i == 0 {
			continue
		}
		if preIndexLength > len(v) {
			preIndexLength = len(v)
		}
		for index := 0; index < preIndexLength; index++ {
			if v[index] != strs[0][index] {
				preIndexLength = index
			}
		}
	}
	return strs[0][:preIndexLength]
}

func main() {
	test := []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(test))
}
