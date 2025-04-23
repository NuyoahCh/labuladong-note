package main

import "github.com/gogo/protobuf/test/indeximport-issue72/index"

type MyArrayList struct {
	// 真正存储数据的底层数组
	data []interface{}
	// 记录当前元素个数
	size int
}

// 设置初始容量
const INIT_CAP = 1

// 创建一个数组链表
func NewMyArrayList() *MyArrayList {
	// 设置初始化数组的容量
	return NewMyArrayListWithCapacity(INIT_CAP)
}

// 创建一个带有容量的数组链表
func NewMyArrayListWithCapacity(initCapacity int) *MyArrayList {
	return &MyArrayList{
		data: make([]interface{}, initCapacity),
		size: 0,
	}
}

//增
func (list *MyArrayList) AddLast(value interface{}) {
	// 计算列表的容量
	cap := len(list.data)
	// 大小和数组的容量相等
	if list.size == cap {
		// 扩容
		list.resize(2 * cap)
	}
	// 在尾部插入数据
	list.data[list.size] = value
	// 增加列表的大小
	list.size++
}

func (list *MyArrayList) Add(index int, value interface{}) error {
	
}
