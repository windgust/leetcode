package main

import (
	"fmt"
	"math"
)

func largestTriangleArea(points [][]int) float64 {
	var maxArea float64
	for i := 0; i < len(points)-2; i++ {
		for j := i + 1; j < len(points)-1; j++ {
			for k := j + 1; k < len(points); k++ {
				s := float64(points[i][0]*points[j][1] + points[j][0]*points[k][1] + points[k][0]*points[i][1] -
					points[i][0]*points[k][1] - points[j][0]*points[i][1] - points[k][0]*points[j][1]) / 2
				s = math.Abs(s)
				if maxArea < s {
					maxArea = s

				}
			}
		}
	}
	return maxArea
}

func main() {
	points := [][]int{{6, 3}, {5, 2}, {5, 8}, {0, 6}}
	fmt.Println(largestTriangleArea(points))
}
