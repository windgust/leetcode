package main

func maxIncreaseKeepingSkyline(grid [][]int) int {
	n := len(grid)
	maxRow, maxCol := make([]int, n, n), make([]int, n, n)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			maxRow[i] = max(maxRow[i], grid[i][j])
			maxCol[i] = max(maxCol[i], grid[j][i])
		}
	}

	var res int
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			res += min(maxRow[i], maxCol[j]) - grid[i][j]
		}
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
