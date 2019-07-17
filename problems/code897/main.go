package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var result = new(TreeNode)

func increasingBST(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	first := result
	getRightTree(root)
	return first.Right
}

func getRightTree(root *TreeNode) {
	if root == nil {
		return
	}
	getRightTree(root.Left)
	result.Right = root
	r := root.Right
	root.Right = nil
	root.Left = nil
	result = result.Right
	getRightTree(r)
}
