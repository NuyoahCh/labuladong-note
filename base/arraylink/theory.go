package main

import "fmt"

func main() {
	// 长度为 5 的数组
	arr := []int{1, 2, 3, 4, 5}
	i := 0
	// 模拟环形数组，这个循环永远不会结束
	for i < len(arr) {
		fmt.Println(arr[i])
		i = (i + 1) % len(arr)
	}
}
