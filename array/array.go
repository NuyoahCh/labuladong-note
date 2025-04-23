package main

import (
	"errors"
	"fmt"
)

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

// 增
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
	// 检查索引是否越界
	if err := list.checkPositionIndex(index); err != nil {
		return err
	}
	// 计算长度
	cap := len(list.data)
	// 查看一下空间是否足够
	if list.size == cap {
		list.resize(2 * cap)
	}

	// 搬移数据 data[index..] -> data[index+1..]
	// 给新元素腾出位置
	for i := list.size - 1; i >= index; i-- {
		list.data[i+1] = list.data[i]
	}
	// 插入新元素
	list.data[index] = value

	// 大小递增
	list.size++
	return nil
}

// 头插法
func (list *MyArrayList) AddFirst(value interface{}) error {
	// 放到第一个元素当中
	return list.Add(0, value)
}

// 删
func (list *MyArrayList) RemoveLast() (interface{}, error) {
	// 删除时候元素的数量不能为0
	if list.size == 0 {
		return nil, errors.New("No such element")
	}
	// 列表数据元素的长度
	cap := len(list.data)

	if list.size == cap/4 {
		// 可以进行缩容，节约空间
		list.resize(cap / 2)
	}

	deletedVal := list.data[list.size-1]
	// 删除最后一个元素
	// 必须给最后一个元素置为 nil，否则就会内存泄露
	list.data[list.size-1] = nil
	list.size--

	return deletedVal, nil
}

// 移除元素
func (list *MyArrayList) Remove(index int) (interface{}, error) {
	// 检查索引越界
	if err := list.checkElementIndex(index); err != nil {
		return nil, err
	}
	// 计算长度
	cap := len(list.data)
	// 进行缩容
	if list.size == cap/4 {
		list.resize(cap / 2)
	}

	// 确定索引的位置，判断删除的值
	deletedVal := list.data[index]

	// 搬移数据 data[index+1..] -> data[index..]
	for i := index + 1; i < list.size; i++ {
		list.data[i-1] = list.data[i]
	}
	// 将最后的元素置为 nil
	// 避免内存的泄漏
	list.data[list.size-1] = nil
	list.size--

	return deletedVal, nil
}

// 删除头位置元素
func (list *MyArrayList) RemoveFirst() (interface{}, error) {
	// 移除索引为0的元素
	return list.Remove(0)
}

// 查
func (list *MyArrayList) Get(index int) (interface{}, error) {
	// 检查索引越界
	if err := list.checkElementIndex(index); err != nil {
		return nil, err
	}
	// 返回要查询的数值的位置
	return list.data[index], nil
}

// 改
func (list *MyArrayList) Set(index int, value interface{}) (interface{}, error) {
	// 检查索引越界
	if err := list.checkElementIndex(index); err != nil {
		return nil, err
	}
	// 修改数据
	oldVal := list.data[index]
	list.data[index] = value
	return oldVal, nil
}

// 工具方法
func (list *MyArrayList) Size() int {
	return list.size
}

func (list *MyArrayList) IsEmpty() bool {
	return list.size == 0
}

// 将 data 的容量改为 newCap
func (list *MyArrayList) resize(newCap int) {
	// 获取初步的长度
	temp := make([]interface{}, newCap)

	// 将之前的数据遍历到 temp 中
	for i := 0; i < list.size; i++ {
		temp[i] = list.data[i]
	}
	list.data = temp
}

func (list *MyArrayList) isElementIndex(index int) bool {
	return index >= 0 && index < list.size

}

func (list *MyArrayList) isPositionIndex(index int) bool {
	return index >= 0 && index <= list.size
}

// 检查 index 索引位置是否可以存在元素
func (list *MyArrayList) checkElementIndex(index int) error {
	if !list.isElementIndex(index) {
		return fmt.Errorf("Index: %d, Size: %d", index, list.size)
	}
	return nil
}

// 检查 index 索引位置是否可以添加元素
func (list *MyArrayList) checkPositionIndex(index int) error {
	if !list.isPositionIndex(index) {
		return fmt.Errorf("Index: %d, Size: %d", index, list.size)
	}
	return nil
}

func (list *MyArrayList) Display() {
	fmt.Printf("size = %d cap = %d\n", list.size, len(list.data))
	fmt.Println(list.data)
}

func main() {
	// 初始容量设为 3
	arr := NewMyArrayListWithCapacity(3)

	// 添加 5 个元素
	for i := 1; i <= 5; i++ {
		arr.AddLast(i)
	}

	arr.Remove(3)
	arr.Add(1, 9)
	arr.AddFirst(100)
	arr.RemoveLast()

	// 100 1 9 2 3
	for i := 0; i < arr.Size(); i++ {
		val, _ := arr.Get(i)
		fmt.Println(val)
	}
}