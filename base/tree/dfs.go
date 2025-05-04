package main

// TreeNode 使用 tree.go 文件当中的
func traverse(root *TreeNode) {
	if root == nil {
		return
	}
	traverse(root.Left)
	traverse(root.Right)
}
