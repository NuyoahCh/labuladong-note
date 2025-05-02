package main

// MyArrayStack 用数组切片作为底层数据结构实现栈
type MyArrayStack[T any] struct {
	list []T
}

// Push 向栈顶加入元素，时间复杂度是 O(1)
func (s *MyArrayStack[T]) Push(e T) {
	s.list = append(s.list, e)
}

// Pop 从栈顶弹出元素，时间复杂度是 O(1)
func (s *MyArrayStack[T]) Pop() T {
	// 如果栈的长度是 0
	if len(s.list) == 0 {
		var zero T
		// 直接返回这个常数
		return zero
	}
	// 获取栈的长度
	e := s.list[len(s.list)-1]
	// 在栈顶进行弹出
	s.list = s.list[:len(s.list)-1]
	return e
}

// Peek 查看栈顶元素，时间复杂度是 O(1)
func (s *MyArrayStack[T]) Peek() T {
	if len(s.list) == 0 {
		var zero T
		return zero
	}
	return s.list[len(s.list)-1]
}

// Size 返回栈中的元素个数
func (s *MyArrayStack[T]) Size() int {
	return len(s.list)
}
