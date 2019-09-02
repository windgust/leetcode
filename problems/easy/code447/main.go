package main

func numberOfBoomerangs(points [][]int) int {
	var cnt int
	for _, p1 := range points {
		set := make(map[int]int)
		for _, p2 := range points {
			dist := (p1[0]-p2[0])*(p1[0]-p2[0]) + (p1[1]-p2[1])*(p1[1]-p2[1])
			set[dist]++
		}
		for _, v := range set {
			cnt += v * (v - 1)
		}
	}
	return cnt
}
