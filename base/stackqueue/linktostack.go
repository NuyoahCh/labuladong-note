package main

import (
	"container/list"
	"fmt"
)

// MyLinkedStack 用链表作为底层数据结构实现栈
type MyLinkedStack struct {
	list *list.List
}

// NewMyLinkedStack 初始化新的链表实现栈
func NewMyLinkedStack() *MyLinkedStack {
	return &MyLinkedStack{list: list.New()}
}

// Push 向栈顶加入元素，时间复杂度 O(1)
func (s *MyLinkedStack) Push(e interface{}) {
	s.list.PushBack(e)
}

// Pop 从栈顶弹出元素，时间复杂度 O(1)
func (s *MyLinkedStack) Pop() interface{} {
	element := s.list.Back()
	if element != nil {
		s.list.Remove(element)
		return element.Value
	}
	return nil
}

// Peek 查看栈顶元素，时间复杂度 O(1)
func (s *MyLinkedStack) Peek() interface{} {
	element := s.list.Back()
	if element != nil {
		return element
	}
	return nil
}

// Size 返回栈中的元素个数，时间复杂度 O(1)
func (s *MyLinkedStack) Size() int {
	return s.list.Len()
}

// test
func main() {
	stack := NewMyLinkedStack()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	fmt.Println(stack.Pop())  // 3
	fmt.Println(stack.Pop())  // 2
	fmt.Println(stack.Peek()) // 1
}
