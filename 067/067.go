package main

import (
	"fmt"
	"strconv"
	"strings"
)

func addBinary(a string, b string) string {
	if len(a) < len(b) {
		a, b = b, a
	}
	addFalg := 0
	var sum []int
	for index := 0; index < len(a); index++ {
		one, err := strconv.Atoi(string(a[len(a)-1-index]))
		if err != nil {
			fmt.Println(err)
		}
		if index <= len(b)-1 {
			two, err := strconv.Atoi(string(b[len(b)-1-index]))
			if err != nil {
				fmt.Println(err)
			}
			if one+two+addFalg >= 2 {
				sum = append(sum, (one+two+addFalg)%2)
				addFalg = (one + two + addFalg) / 2
			} else {
				sum = append(sum, one+two+addFalg)
				addFalg = 0
			}
		} else {
			if one+addFalg >= 2 {
				sum = append(sum, (one+addFalg)%2)
				addFalg = (one + addFalg) / 2
			} else {
				sum = append(sum, one+addFalg)
				addFalg = 0
			}
		}
	}
	if addFalg == 1 {
		sum = append(sum, 1)
	}
	for i, j := 0, len(sum)-1; i < j; i, j = i+1, j-1 {
		sum[i], sum[j] = sum[j], sum[i]
	}
	return strings.Trim(strings.Replace(fmt.Sprint(sum), " ", "", -1), "[]")
}
func main() {
	fmt.Println(addBinary("101111", "10"))
}
