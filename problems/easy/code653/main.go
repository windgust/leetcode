package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTarget(root *TreeNode, k int) bool {
	m := make(map[int]struct{})
	return findTargetByMap(root, k, m)

}

func findTargetByMap(node *TreeNode, k int, m map[int]struct{}) bool {
	if node == nil {
		return false
	}
	_, ok := m[k-node.Val]
	m[node.Val] = struct{}{}
	return ok || findTargetByMap(node.Left, k, m) || findTargetByMap(node.Right, k, m)
}
