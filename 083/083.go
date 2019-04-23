package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	t := head
	for head != nil && head.Next != nil {
		if head.Val == head.Next.Val {
			temp := head.Next
			for temp.Next != nil && temp.Val == temp.Next.Val {
				temp = temp.Next
			}
			if temp.Next != nil {
				head.Next = temp.Next
			} else {
				head.Next = nil
			}

		}
		head = head.Next
	}
	return t
}
func main() {
	test := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val:  1,
			Next: nil,
		},
	}
	result := deleteDuplicates(test)
	for result != nil {
		fmt.Println(result.Val)
		result = result.Next
	}
}
