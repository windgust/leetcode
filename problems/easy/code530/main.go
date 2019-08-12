package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getMinimumDifference(root *TreeNode) int {
	pre := -1
	min := math.MaxInt32
	minDiff(root, &pre, &min)
	return min
}

func minDiff(root *TreeNode, pre, min *int) {
	if root == nil {
		return
	}
	minDiff(root.Left, pre, min)
	if *pre != -1 {
		*min = Min(abs(root.Val-*pre), *min)
	}
	*pre = root.Val
	minDiff(root.Right, pre, min)
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}
