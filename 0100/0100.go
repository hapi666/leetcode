package main

import "fmt"

// TreeNode is a tree struct
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if q == nil && p == nil {
		return true
	}
	if q == nil || p == nil {
		return false
	}
	if p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right) {
		return true
	}
	return false
}

func main() {
	t1 := &TreeNode{
		Val:  2,
		Left: nil,
		Right: &TreeNode{
			Val:   0,
			Left:  nil,
			Right: nil,
		},
	}
	t2 := &TreeNode{
		Val: 0,
		Left: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
		Right: nil,
	}
	fmt.Println(isSameTree(t1, t2))
}
