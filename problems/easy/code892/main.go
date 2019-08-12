package main

import "fmt"

func surfaceArea(grid [][]int) int {
	var sum int
	n := len(grid)
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 0 {
				continue
			}
			sum += (grid[i][j] << 2) + 2
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n-1; j++ {
			m := min(grid[i][j], grid[i][j+1])
			m += min(grid[j][i], grid[j+1][i])
			sum -= m << 1
		}
	}
	return sum
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	grid := [][]int{{2, 2, 2}, {2, 1, 2}, {2, 2, 2}}
	fmt.Println(surfaceArea(grid))
}
