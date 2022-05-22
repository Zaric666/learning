package binarytree

import (
	"fmt"
	"testing"
)

// 前序遍历，根左右
func preOrderTraverse(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Println(root.Data)
	preOrderTraverse(root.Left)
	preOrderTraverse(root.Right)
}

//中序遍历，左根右
func inOrderTraverse(root *TreeNode) {
	if root == nil {
		return
	}
	preOrderTraverse(root.Left)
	fmt.Println(root.Data)
	preOrderTraverse(root.Right)
}

// 后序遍历，左右根
func postOrderTraverse(root *TreeNode) {
	if root == nil {
		return
	}
	preOrderTraverse(root.Left)
	preOrderTraverse(root.Right)
	fmt.Println(root.Data)
}

//	  5
//   / \
//  2   7
// 前序 5->2->7
// 中序 2->5->7
// 后序 2->7->5
func TestTraversal(t *testing.T) {
	var root *TreeNode
	root = &TreeNode{5, &TreeNode{2, nil, nil}, &TreeNode{7, nil, nil}}
	fmt.Println("========前序======")
	preOrderTraverse(root)
	fmt.Println("========中序======")
	inOrderTraverse(root)
	fmt.Println("========后序======")
	postOrderTraverse(root)
}
