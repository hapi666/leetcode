package main

import "fmt"

var runesMap = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func romanToInt(s string) int {
	result := 0
	index := 0
	l := len(s)
	for index = 0; index < l-1; {
		if runesMap[string(s[index])] < runesMap[string(s[index+1])] {
			result += runesMap[string(s[index+1])] - runesMap[string(s[index])]
			index += 2
		} else {
			result += runesMap[string(s[index])]
			index++
		}
	}
	if index == l-1 {
		result += runesMap[string(s[l-1])]
	}
	return result
}

func main() {
	str1 := "MCMXCIV"
	str2 := "IV"
	str3 := "IX"
	str4 := "LVIII"
	fmt.Println(romanToInt(str1))
	fmt.Println(romanToInt(str2))
	fmt.Println(romanToInt(str3))
	fmt.Println(romanToInt(str4))
}
