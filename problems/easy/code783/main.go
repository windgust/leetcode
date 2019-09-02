package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDiffInBST(root *TreeNode) int {
	pre := -1
	min := math.MaxInt32
	minDiff(root, &pre, &min)
	return min
}

func minDiff(node *TreeNode, pre, min *int) {
	if node == nil {
		return
	}
	minDiff(node.Left, pre, min)
	if *pre != -1 {
		*min = Min(*min, node.Val-*pre)
	}
	*pre = node.Val
	minDiff(node.Right, pre, min)

}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
