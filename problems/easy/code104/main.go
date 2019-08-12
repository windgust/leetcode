package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var depth int
	maxDep := 1
	searcgMaxDepth(root, depth, &maxDep)
	return maxDep
}

func searcgMaxDepth(root *TreeNode, depth int, maxDep *int) {
	if root == nil {
		if *maxDep < depth {
			*maxDep = depth
		}
		return
	}
	searcgMaxDepth(root.Left, depth+1, maxDep)
	searcgMaxDepth(root.Right, depth+1, maxDep)
}
