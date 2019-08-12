package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func tree2str(t *TreeNode) string {
	if t == nil {
		return ""
	}
	if t.Left == nil && t.Right == nil {
		return fmt.Sprintf("%d", t.Val)
	}
	if t.Right == nil {
		return fmt.Sprintf("%d(%s)", t.Val, tree2str(t.Left))
	}
	return fmt.Sprintf("%d(%s)(%s)", t.Val, tree2str(t.Left), tree2str(t.Right))
}
