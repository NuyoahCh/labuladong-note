package main

import "fmt"

func main() {
	// 大小为 10 的数组里面已经装上了 4 个元素
	var arr [10]int
	for i := 0; i < 4; i++ {
		arr[i] = i
	}
	// 在索引 2 插入元素 666
	// 需要把索引 2 以及之后的元素都往后面进行移位，移动一位
	// 注意要倒着遍历数组中已有的元素，避免覆盖
	for i := 4; i > 2; i-- {
		arr[i] = arr[i-1]
	}

	arr[2] = 666
	fmt.Println(arr)
}
