package main

import "fmt"

func sumEvenAfterQueries(A []int, queries [][]int) []int {
	var sum int
	var res []int
	for _, a := range A {
		if (a & 1) == 0 {
			sum += a
		}
	}

	for i := range queries {
		val, index := queries[i][0], queries[i][1]
		tmp := A[index] + val
		if (A[index] & 1) == 0 {
			sum -= A[index]
		}
		if (tmp & 1) == 0 {
			sum += tmp
		}

		A[index] = tmp
		res = append(res, sum)
	}
	return res
}

func main() {
	A := []int{1, 2, 0, 4}
	queries := [][]int{{5,2}, {0,2}, {3,3}, {2,1}}
	fmt.Println(sumEvenAfterQueries(A, queries))
}
