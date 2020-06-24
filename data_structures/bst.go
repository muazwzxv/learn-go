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
	node := &node{data, nil, nil}
	return node
}

func printInOrder(node *node) {
	if node == nil {
		return
	}
	printInOrder(node.left)
	fmt.Println(node.data, " ")
	printInOrder(node.right)
}

func printPreOrder(node *node) {
	if node == nil {
		return
	}
	fmt.Println(node.data, " ")
	printPreOrder(node.left)
	printPreOrder(node.right)
}

func printPostOrder(node *node) {
	if node == nil {
		return
	}
	printPostOrder(node.left)
	printPostOrder(node.left)
	fmt.Println(node.data, " ")
}
