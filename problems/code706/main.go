package main

import "fmt"

type kv struct {
	key   int
	value int
}

type MyHashMap struct {
	v   [][]*kv
	cnt int
}

/** Initialize your data structure here. */
func Constructor() MyHashMap {

	return MyHashMap{
		v: make([][]*kv, 8192),
	}
}

/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int) {
	this.putValue(key, value, this.v)

	if this.cnt > int(float64(len(this.v))*0.75) {
		tmp := make([][]*kv, len(this.v)*2)
		for i := range this.v {
			if this.v[i] == nil {
				continue
			}
			for j := range this.v[i] {
				this.putValue(this.v[i][j].key, this.v[i][j].value, tmp)
			}
		}
		this.v = tmp
	}

}

/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
	h := key & (len(this.v) - 1)
	if this.v[h] == nil {
		return -1
	}
	kvs := this.v[h]
	for i := range kvs {
		if kvs[i].key == key {
			return kvs[i].value
		}
	}
	return -1
}

/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int) {
	h := key & (len(this.v) - 1)
	if this.v[h] == nil {
		return
	}
	kvs := this.v[h]
	for i := range kvs {
		if kvs[i].key == key {
			this.v[h] = append(kvs[:i], kvs[i+1:]...)
			this.cnt--
			break
		}
	}
}

func (this *MyHashMap) putValue(key, value int, m [][]*kv) {
	h := key & (len(m) - 1)
	if m[h] == nil {
		m[h] = []*kv{
			&kv{
				key:   key,
				value: value,
			},
		}
		this.cnt++
		return
	}

	kvs := m[h]
	for i := range kvs {
		if kvs[i].key == key {
			kvs[i].value = value
			return
		}
	}
	m[h] = append(m[h], &kv{
		key:   key,
		value: value,
	})
	this.cnt++
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */



func main() {
	hashMap := Constructor()
	hashMap.Put(16, 1);
	hashMap.Put(33, 2);
	hashMap.Put(34, 2);

	fmt.Println(hashMap.Get(16)); // returns 1
	fmt.Println(hashMap.Get(33)); // returns -1 (not found)
	hashMap.Put(49, 12);           // update the existing value
	fmt.Println(hashMap.Get(49)); // returns 1
	hashMap.Remove(16);           // remove the mapping for 2
	fmt.Println(hashMap.Get(16)); // returns -1 (not found)
}
