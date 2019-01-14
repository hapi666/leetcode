package main

import "fmt"

var (
	max = 1<<31 - 1
	min = -1 << 31
)

func reverse(x int) int {
	result := 0
	for x != 0 {
		result = result*10 + x%10
		x = x / 10
	}
	if result > max || result < min {
		return 0
	}
	return result
}

func main() {
	data := 100
	r := reverse(data)
	fmt.Println(r)
}
