package main

import "fmt"

type State struct {
	node  *TreeNode
	depth int
}

func levelOrderTraverse1(root *TreeNode) {
	if root == nil {
		return
	}
	q := []*TreeNode{}
	q = append(q, root)
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		// 访问 cur 节点
		fmt.Println(cur.Val)

		// 把 cur 的左右子节点加入到队列中
		if cur.Left != nil {
			q = append(q, cur.Left)
		}
		if cur.Right != nil {
			q = append(q, cur.Right)
		}
	}
}

func levelOrderTraverse2(root *TreeNode) {
	if root == nil {
		return
	}
	q := []*TreeNode{root}
	depth := 1
	for len(q) > 0 {

		sz := len(q)
		for i := 0; i < sz; i++ {
			cur := q[0]
			q = q[1:]
			fmt.Printf("depth = %d, val = %d\n", depth, cur.Val)

			if cur.Left != nil {
				q = append(q, cur.Left)
			}
			if cur.Right != nil {
				q = append(q, cur.Right)
			}
		}
		depth++
	}
}

func levelOrderTraverse3(root *TreeNode) {
	if root == nil {
		return
	}
	q := []State{{root, 1}}

	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		fmt.Printf("depth = %d, val = %d\n", cur.depth, cur.node.Val)

		if cur.node.Left != nil {
			q = append(q, State{cur.node.Left, cur.depth + 1})
		}
		if cur.node.Right != nil {
			q = append(q, State{cur.node.Right, cur.depth + 1})
		}
	}
}
