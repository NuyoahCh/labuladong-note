package main

// MyArrayQueue 数组实现队列
type MyArrayQueue[E any] struct {
	arr *CycleArray[E]
}

// NewMyArrayQueue 自定义数组队列
func NewMyArrayQueue[E any]() *MyArrayQueue[E] {
	return &MyArrayQueue[E]{
		arr: NewCircleArray[E](),
	}
}

// Push 添加元素
func (q *MyArrayQueue[E]) Push(t E) {
	q.arr.AddLast(t)
}

// Pop 移除首部元素
func (q *MyArrayQueue[E]) Pop() error {
	return q.arr.RemoveFirst()
}

// Peek 查看第一个元素
func (q *MyArrayQueue[E]) Peek() (E, interface{}) {
	return q.arr.GetFirst()
}

// Size 获取数组队列的长度
func (q *MyArrayQueue[E]) Size() int {
	return q.arr.Size()
}
