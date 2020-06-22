package main

import (
	"errors"
	"fmt"
)

type node struct {
	data *int
	next *node
}

type linkedlist struct {
	head *node
	last *node
}

func newNode(data int) *node {
	return &node{&data, nil}
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

func displayList(list *linkedlist) error {
	current := list.head

	if current.data == nil {
		return errors.New("The list is empty")
	}

	for current.next != nil {
		fmt.Println(current.data)
		current = current.next
	}

	return nil
}

func main() {
	fmt.Println("LinkedList")
}
