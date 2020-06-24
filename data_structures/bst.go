package main

import "fmt"

type node struct {
	data  int
	left  *node
	right *node
}

type tree struct {
	root *node
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func newNode(data int) *node {
	return &node{data, nil, nil}
}

func printInOrder(node *node) {
	if node == nil {
		return
	}
	printInOrder(node.left)
	fmt.Println(node.data, " ")
	printInOrder(node.right)
}

func insert(root *node, data int) {
	if root == nil {
		newNode(data)
	}

	if data < root.data {
		insert(root.left, data)
	} else {
		insert(root.right, data)
	}
}

// return the furthest left element
func inOrderSuccessor(root *node) *node {
	current := root
	if current.left != nil {
		current = current.left
	}
	return current
}

func delete(root *node, toDelete int) *node {
	if root == nil {
		return nil
	}

	if toDelete > root.data {
		root.right = delete(root.right, toDelete)
	} else if toDelete < root.data {
		root.left = delete(root.left, toDelete)
	} else {
		// Reach the point to delete nodes

		if root.left == nil {
			root = root.right
		} else if root.right == nil {
			root = root.left
		} else {

		}
	}

}
