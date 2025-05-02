package main

type MyArrayDeque[E any] struct {
	arr CycleArray[E]
}

// NewMyArrayDeque creates a new MyArrayDeque
func NewMyArrayDeque[E any]() *MyArrayDeque[E] {
	return &MyArrayDeque[E]{arr: CycleArray[E]{}}
}

// AddFirst 从队头插入元素，时间复杂度 O(1)
func (d *MyArrayDeque[E]) AddFirst(e E) {
	d.arr.AddFirst(e)
}

// AddLast 从队尾插入元素，时间复杂度 O(1)
func (d *MyArrayDeque[E]) AddLast(e E) {
	d.arr.AddLast(e)
}

// RemoveFirst 从队头删除元素，时间复杂度 O(1)
func (d *MyArrayDeque[E]) RemoveFirst() interface{} {
	return d.arr.RemoveFirst()
}

// RemoveLast 从队尾删除元素，时间复杂度 O(1)
func (d *MyArrayDeque[E]) RemoveLast() interface{} {
	return d.arr.RemoveLast()
}

// PeekFirst 查看队头元素，时间复杂度 O(1)
func (d *MyArrayDeque[E]) PeekFirst() (E, interface{}) {
	return d.arr.GetFirst()
}

// PeekLast 查看队尾元素，时间复杂度 O(1)
func (d *MyArrayDeque[E]) PeekLast() (E, interface{}) {
	return d.arr.GetLast()
}
