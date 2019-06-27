package main

import "fmt"

type RecentCounter struct {
	counter []int
}


func Constructor() RecentCounter {
	return RecentCounter{
		counter:make([]int, 0, 0),
	}
}


func (this *RecentCounter) Ping(t int) int {
	this.counter = append(this.counter, t)
	k := 0
	for _, i := range this.counter {
		if i >= t - 3000 {
			break
		}
		k++
	}
	this.counter = this.counter[k:]
	return len(this.counter)
}

func main() {
	obj := Constructor()
	fmt.Println(obj.Ping(1))
	fmt.Println(obj.Ping(100))
	fmt.Println(obj.Ping(3001))
	fmt.Println(obj.Ping(3002))

}


/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */
