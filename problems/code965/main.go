package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isUnivalTree(root *TreeNode) bool {
	reslut := true
	if root == nil {
		return reslut
	}
	if root.Left != nil  {
		reslut = reslut && root.Val == root.Left.Val
	}
	if root.Right != nil {
		reslut = reslut && root.Val == root.Right.Val
	}
	if !reslut {
		return reslut
	}
	return reslut && isUnivalTree(root.Left) && isUnivalTree(root.Right)
}
