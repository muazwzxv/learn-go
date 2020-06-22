package main

import (
	"fmt"
)

type node struct {
	data int
	next *node
}

type linkedlist struct {
	head *node
	last *node
}

func newNode(data int) *node {
	return &node{data, nil}
}

func (list *linkedlist) addAtFront(data int) {
	node := newNode(data)
	node.next = list.head
	list.head = node
}

func (list *linkedlist) addAtBack(data int) {
	node := newNode(data)

	if list.head == nil {
		list.head = node
	}
	current := list.head

	for current.next != nil {
		current = current.next
	}

	current.next = node
}

func main() {
	fmt.Println("LinkedList")
}
