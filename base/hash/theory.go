package main

// MyHashMap 哈希表的伪代码实现
type MyHashMap struct {
	table []interface{}
}

// Constructor 构造器
func Constructor() MyHashMap {
	return MyHashMap{
		table: make([]interface{}, 10000),
	}
}

// Put 增/改，复杂度 O(1)
func (h *MyHashMap) Put(key, value interface{}) {
	index := h.hash(key)
	h.table[index] = value
}

// Get 查，复杂度 O(1)
func (h *MyHashMap) Get(key, value interface{}) interface{} {
	index := h.hash(key)
	return h.table[index]
}

// Remove 删，复杂度 O(1)
func (h *MyHashMap) Remove(key interface{}) {
	index := h.hash(key)
	h.table[index] = nil
}

// 哈希函数，把 key 转化成 table 中的合法索引
// 时间复杂度必须是 O(1)，才能保证上述方法的复杂度都是 O(1)
func (h *MyHashMap) hash(key interface{}) int {
	return 0
}
