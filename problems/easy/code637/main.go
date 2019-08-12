package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfLevels(root *TreeNode) []float64 {

	var nodes []*TreeNode
	var res []float64
	var i, next, offset, cnt int
	var sum float64
	nodes = append(nodes, root)
	for i < len(nodes) {
		sum += float64(nodes[i].Val)
		cnt++

		if nodes[i].Left != nil {
			nodes = append(nodes, nodes[i].Left)
			offset++
		}
		if nodes[i].Right != nil {
			nodes = append(nodes, nodes[i].Right)
			offset++
		}
		if i == next {
			res = append(res, sum/float64(cnt))
			next += offset
			offset = 0
			sum = 0
			cnt = 0
		}
		i++
	}
	return res
}
