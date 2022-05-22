package binarytree

import (
	"fmt"
	"testing"
)

func traverseInvertTree(root *TreeNode) {
	if root == nil {
		return
	}
	root.Right, root.Left = root.Left, root.Right

	traverseInvertTree(root.Left)
	traverseInvertTree(root.Right)
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	left := invertTree(root.Left)
	right := invertTree(root.Right)

	root.Left = right
	root.Right = left

	return root
}

func TestInvertTree(t *testing.T) {
	var root *TreeNode
	root = &TreeNode{
		4,
		&TreeNode{2, &TreeNode{1, nil, nil}, &TreeNode{3, nil, nil}},
		&TreeNode{6, &TreeNode{5, nil, nil}, &TreeNode{7, nil, nil}},
	}
	fmt.Println("=======递归翻转左右子树=======")
	root = invertTree(root)
	inOrderTraverse(root)

	root = &TreeNode{
		4,
		&TreeNode{2, &TreeNode{1, nil, nil}, &TreeNode{3, nil, nil}},
		&TreeNode{6, &TreeNode{5, nil, nil}, &TreeNode{7, nil, nil}},
	}
	fmt.Println("=======遍历翻转=======")
	traverseInvertTree(root)
	inOrderTraverse(root)
}
