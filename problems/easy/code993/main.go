package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type DepthAndParent struct {
	Depth  int
	Parent int
}

func isCousins(root *TreeNode, x int, y int) bool {
	depthAndParent := make(map[int]*DepthAndParent)
	findDepthAndParent(root, &TreeNode{Val: -1}, x, y, 0, depthAndParent)
	return depthAndParent[x].Depth == depthAndParent[y].Depth && depthAndParent[x].Parent != depthAndParent[y].Parent
}

func findDepthAndParent(root, parentNode *TreeNode, x, y, depth int, depthAndParent map[int]*DepthAndParent) {
	if root == nil {
		return
	}
	if root.Val == x || root.Val == y {
		depthAndParent[root.Val] = &DepthAndParent{
			Depth:  depth,
			Parent: parentNode.Val,
		}
	}
	if len(depthAndParent) < 2 {
		findDepthAndParent(root.Left, root, x, y, depth+1, depthAndParent)
	}
	if len(depthAndParent) < 2 {
		findDepthAndParent(root.Right, root, x, y, depth+1, depthAndParent)
	}
}
