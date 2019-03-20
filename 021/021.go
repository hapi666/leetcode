package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func insertNode(l *ListNode, key int) {
	for l != nil {
		if l.Val <= key {
			temp := new(ListNode)
			temp.Val = key
			if l.Next == nil {
				l.Next = temp
				break
			} else if l.Next.Val > key {
				temp.Next = l.Next
				l.Next = temp
				break
			}
		}
		l = l.Next
	}
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if nil == l1 {
		return l2
	}
	if nil == l2 {
		return l1
	}
	if l1.Val > l2.Val {
		for l1 != nil {
			insertNode(l2, l1.Val)
			l1 = l1.Next
		}
		return l2
	} else {
		for l2 != nil {
			insertNode(l1, l2.Val)
			l2 = l2.Next
		}
		return l1
	}

	//if nil == l1 {
	//	return l2
	//}
	//if nil == l2 {
	//	return l1
	//}
	//
	//if l1.Val < l2.Val {
	//	l1.Next = mergeTwoLists(l1.Next, l2)
	//	return l1
	//} else {
	//	l2.Next = mergeTwoLists(l1, l2.Next)
	//	return l2
	//}
}

func main() {
	test1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	test2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	t := mergeTwoLists(test1, test2)
	for t != nil {
		fmt.Println(t.Val)
		t = t.Next
	}

}
