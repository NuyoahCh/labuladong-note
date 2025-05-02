package main

import (
	"fmt"
	"hash/fnv"
)

type KVNode1[K comparable, V any] struct {
	key K
	val V
}

// MyLinearProbingMap1 线性探查哈希表
type MyLinearProbingMap1[K comparable, V any] struct {
	table []*KVNode1[K, V]
	size  int
}

const InitCap = 4

func NewMyLinearProbingMap1[K comparable, V any]() *MyLinearProbingMap1[K, V] {
	return NewMyLinearProbingMap1WithCapacity[K, V](InitCap)
}

func NewMyLinearProbingMap1WithCapacity[K comparable, V any](initCapacity int) *MyLinearProbingMap1[K, V] {
	return &MyLinearProbingMap1[K, V]{
		table: make([]*KVNode1[K, V], initCapacity),
	}
}

func (m *MyLinearProbingMap1[K, V]) Put(key K, val V) error {
	if m.size >= len(m.table)*3/4 {
		m.resize(len(m.table) * 2)
	}

	index := m.getKeyIndex(key)
	if m.table[index] != nil && m.table[index].key == key {
		m.table[index].val = val
		return nil
	}

	m.table[index] = &KVNode1[K, V]{key, val}
	m.size++
	return nil
}

func (m *MyLinearProbingMap1[K, V]) Remove(key K) error {
	if m.size <= len(m.table)/8 {
		m.resize(len(m.table) / 2)
	}

	index := m.getKeyIndex(key)
	if m.table[index] == nil {
		return nil
	}

	m.table[index] = nil
	m.size--
	index = (index + 1) % len(m.table)
	for m.table[index] != nil {
		entry := m.table[index]
		m.table[index] = nil
		m.size--
		m.Put(entry.key, entry.val)
		index = (index + 1) % len(m.table)
	}
	return nil
}

func (m *MyLinearProbingMap1[K, V]) Get(key K) (V, error) {
	index := m.getKeyIndex(key)
	if m.table[index] == nil {
		return *new(V), nil
	}
	return m.table[index].val, nil
}

func (m *MyLinearProbingMap1[K, V]) ContainsKey(key K) bool {
	index := m.getKeyIndex(key)
	return m.table[index] != nil
}

func (m *MyLinearProbingMap1[K, V]) Keys() []K {
	keys := []K{}
	for _, entry := range m.table {
		if entry != nil {
			keys = append(keys, entry.key)
		}
	}
	return keys
}

/**
其他的工具函数
*/

func (m *MyLinearProbingMap1[K, V]) Size() int {
	return m.size
}

func (m *MyLinearProbingMap1[K, V]) hash(key K) int {
	h := fnv.New32a()
	h.Write([]byte(fmt.Sprintf("%v", key)))
	return int(h.Sum32()) & 0x7fffffff % len(m.table)
}

func (m *MyLinearProbingMap1[K, V]) getKeyIndex(key K) int {
	index := m.hash(key)
	for m.table[index] != nil {
		if m.table[index].key == key {
			return index
		}
		index = (index + 1) % len(m.table)
	}
	return index
}

func (m *MyLinearProbingMap1[K, V]) resize(newCap int) {
	newMap := NewMyLinearProbingMap1WithCapacity[K, V](newCap)
	for _, entry := range m.table {
		if entry != nil {
			newMap.Put(entry.key, entry.val)
		}
	}
	m.table = newMap.table
}


func main() {
	mapInstance := NewMyLinearProbingMap1[int, int]()
	mapInstance.Put(1, 1)
	mapInstance.Put(2, 2)
	mapInstance.Put(10, 10)
	mapInstance.Put(20, 20)
	mapInstance.Put(30, 30)
	mapInstance.Put(3, 3)

	val, _ := mapInstance.Get(1)
	fmt.Println(val) // 1
	val, _ = mapInstance.Get(2)
	fmt.Println(val) // 2
	val, _ = mapInstance.Get(20)
	fmt.Println(val) // 20

	mapInstance.Put(1, 100)
	val, _ = mapInstance.Get(1)
	fmt.Println(val) // 100

	mapInstance.Remove(20)
	fmt.Println(mapInstance.ContainsKey(20)) // false
	val, _ = mapInstance.Get(30)
	fmt.Println(val) // 30
}