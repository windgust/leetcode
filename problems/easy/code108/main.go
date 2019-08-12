package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	return bst(nums, 0, len(nums)-1)
}

func bst(nums []int, l, r int) *TreeNode {
	if l > r {
		return nil
	}

	mid := (l + r) >> 1
	return &TreeNode{
		Val:   nums[mid],
		Left:  bst(nums, l, mid-1),
		Right: bst(nums, mid+1, r),
	}
}
