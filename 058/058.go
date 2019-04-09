package main

import "fmt"

func lengthOfLastWord(s string) int {
	count := 0
	flag := false
	for index := len(s) - 1; index >= 0; index-- {
		if s[index] != ' ' {
			count++
			flag = true
		} else {
			if flag {
				break
			}
		}
	}
	return count
}

func main() {
	s := "Hello World "
	fmt.Println(lengthOfLastWord(s))
}
