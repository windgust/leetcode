package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumRootToLeaf(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var sum int
	sumLeaf(root, 0, &sum)
	return sum
}

func sumLeaf(node *TreeNode, subSum int, sum *int) {
	subSum = (subSum << 1) + node.Val
	if node.Left == nil && node.Right == nil {
		*sum += subSum
		return
	}
	if node.Left != nil {
		sumLeaf(node.Left, subSum, sum)
	}
	if node.Right != nil {
		sumLeaf(node.Right, subSum, sum)
	}
}
