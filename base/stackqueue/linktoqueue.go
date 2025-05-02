package main

import (
	"container/list"
	"fmt"
)

// MyLinkedQueue 用链表作为底层数据结构实现队列
type MyLinkedQueue struct {
	list *list.List
}

// NewMyLinkedQueue 构造函数
func NewMyLinkedQueue() *MyLinkedQueue {
	return &MyLinkedQueue{list: list.New()}
}

// Push 向队尾插入元素，时间复杂度 O(1)
func (q *MyLinkedQueue) Push(e interface{}) {
	q.list.PushBack(e)
}

// Pop 从队头删除元素，时间复杂度 O(1)
func (q *MyLinkedQueue) Pop() interface{} {
	front := q.list.Front()
	if front != nil {
		return q.list.Remove(front)
	}
	return nil
}

// Peek 查看队头元素，时间复杂度 O(1)
func (q *MyLinkedQueue) Peek() interface{} {
	front := q.list.Front()
	if front != nil {
		return front.Value
	}
	return nil
}

// 返回队列中的元素个数，时间复杂度 O(1)
func (q *MyLinkedQueue) Size() int {
	return q.list.Len()
}

func main() {
	queue := NewMyLinkedQueue()
	queue.Push(1)
	queue.Push(2)
	queue.Push(3)
	fmt.Println(queue.Peek()) // 1
	fmt.Println(queue.Pop())  // 1
	fmt.Println(queue.Pop())  // 2
	fmt.Println(queue.Peek()) // 3
}
