package main

import (
	"fmt"
)

const Base = 8
const Offset = 3

type MyHashSet struct {
	Set []uint8
}

/** Initialize your data structure here. */
func Constructor() MyHashSet {
	return MyHashSet{
		Set: make([]uint8, 125000+1),
	}
}

func (this *MyHashSet) Add(key int) {
	if key == 0 {
		this.Set[0] |= 1
		return
	}
	offset := key >> Offset
	this.Set[offset] |= uint8(1 << (uint(key) & (Base - 1)))
}

func (this *MyHashSet) Remove(key int) {
	offset := key >> Offset
	this.Set[offset] &= ^uint8(1 << (uint(key) & (Base - 1)))
}

/** Returns true if this set contains the specified element */
func (this *MyHashSet) Contains(key int) bool {
	offset := key >> Offset
	return (this.Set[offset] & uint8(1<<(uint(key)&(Base-1)))) != 0
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */

func main() {
	obj := Constructor();
	for i := 1; i <= 16; i++ {
		fmt.Println(i, obj.Contains(i))
		obj.Add(i)
		fmt.Println(obj.Contains(i))
		//obj.Remove(i)
		//fmt.Println(obj.Contains(i))
		fmt.Println("*********")
	}
}
