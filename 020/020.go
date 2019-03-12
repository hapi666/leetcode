package main

import (
	"fmt"
	"log"
)

type stack []rune

func (s *stack) push(d rune) {
	*s = append(*s, d)
}

func (s *stack) pop() rune {
	if s == nil {
		log.Println("空栈 无法弹出")
		return ' '
	}
	v := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return v
}

func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}
	if len(s)%2 != 0 {
		return false
	}
	if s[0] == ')' || s[0] == ']' || s[0] == '}' {
		return false
	}
	stack := new(stack)
	for index := 0; index < len(s); index++ {
		switch s[index] {
		case '(':
			stack.push('(')
			break
		case '[':
			stack.push('[')
			break
		case '{':
			stack.push('{')
			break
		case ')':
			if stack.pop() != '(' {
				return false
			}
			break
		case ']':
			if stack.pop() != '[' {
				return false
			}
			break
		case '}':
			if stack.pop() != '{' {
				return false
			}
			break
		}
	}
	if len(*stack) != 0 {
		return false
	}
	return true
}

func main() {
	if isValid("(([])){}") {
		fmt.Println("yes!")
	} else {
		fmt.Println("false!")
	}
}
