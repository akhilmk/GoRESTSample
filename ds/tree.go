package main

import (
	"fmt"
)

func main() {
	fmt.Println("main start")

	node1 := &treeNode{7, nil, nil}
	node2 := &treeNode{12, nil, nil}

	root := treeNode{data: 1}
	root.left = &treeNode{2, node1, nil}
	root.right = &treeNode{3, nil, node2}

	found := findItem(&root, 12)
	fmt.Println(found)
	fmt.Println("main end")
}

// Pass root node and item to find
// DFS, Inorder traversal will be performed on tree
func findItem(node *treeNode, item int) bool {

	if node.data == item {
		return true
	}

	found := false
	if node.left != nil {
		found = findItem(node.left, item)
	}
	if node.right != nil && !found {
		found = findItem(node.right, item)
	}

	return found
}

type treeNode struct {
	data  int
	left  *treeNode
	right *treeNode
}
