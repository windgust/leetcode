package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func convertBST(root *TreeNode) *TreeNode {
	sumGreaterTree(root,0)
	return root
}

func sumGreaterTree(node *TreeNode, pVal int) int {
	if node == nil {
		return pVal
	}
	node.Val += sumGreaterTree(node.Right, pVal)
	return sumGreaterTree(node.Left, node.Val)
}
