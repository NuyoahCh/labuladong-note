package main

import "fmt"

type Node struct {
	val      int
	children []*Node
}

func traverseNary(root *Node) {
	if root == nil {
		return
	}
	for _, child := range root.children {
		traverseNary(child)
	}
}

func levelOrderTraverse4(root *Node) {
	if root == nil {
		return
	}
	q := []*Node{}
	q = append(q, root)
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		fmt.Println(cur.val)

		for _, child := range cur.children {
			q = append(q, child)
		}
	}
}

func levelOrderTraverse5(root *Node) {
	if root == nil {
		return
	}
	q := []*Node{root}
	depth := 1

	for len(q) > 0 {
		sz := len(q)
		for i := 0; i < sz; i++ {
			cur := q[0]
			q = q[1:]
			fmt.Printf("depth = %d, val = %d\n", depth, cur.val)

			for _, child := range cur.children {
				q = append(q, child)
			}
		}
		depth++
	}
}

//type State1 struct {
//	node *Node
//	depth int
//}
//
//func levelOrderTraverse6(root *Node) {
//	if root == nil {
//		return
//	}
//	q := []State{}
//	q=append(q, State1{root, 1})
//
//	for len(q) > 0 {
//		state := q[0]
//		q = q[1:]
//		cur := state.node
//		depth := state.depth
//		fmt.Printf("depth = %d, val = %d\n", depth, cur.val)
//
//		for _, child := range
//	}
//}