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

func insert(root *node, data int) *node {
	if root == nil {
		return newNode(data)
	}

	if data < root.data {
		root.left = insert(root.left, data)
	} else {
		root.right = insert(root.right, data)
	}

	return root
}

// return the furthest left element
func inOrderSuccessor(root *node) *node {
	current := root
	if current.left != nil {
		current = current.left
	}
	return current
}
