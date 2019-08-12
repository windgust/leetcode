package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func bstToGst(root *TreeNode) *TreeNode {
	resetVal(root, 0)
	return root
}

func resetVal(root *TreeNode, pVal int) int {
	if root == nil {
		return pVal
	}
	root.Val = root.Val + resetVal(root.Right, pVal)
	return resetVal(root.Left, root.Val)
}
