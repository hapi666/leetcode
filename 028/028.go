package main

import (
	"fmt"
	"strings"
)

func strStr(haystack string, needle string) int {
	if len(haystack) < len(needle) {
		return -1
	}
	if needle == "" {
		return 0
	}
	return strings.Index(haystack, needle)
}

func main() {
	fmt.Println(strStr("aaaaa", "bb"))
}
